package main

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/common-tools-haonan/docker/cgroup"
	"github.com/common-tools-haonan/docker/cgroup/subsystem"
	"github.com/common-tools-haonan/docker/container"
	"github.com/common-tools-haonan/docker/network"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func fork(isStd bool, image, containerId, volume string, env []string) (cmds *exec.Cmd, write *os.File) {

	read, write, err := os.Pipe()
	if err != nil {
		logrus.Fatal("the process of creating a pipe failed occurring fork, err:%s ", err)
	}
	initSymbol, _ := os.Readlink("/proc/self/exe")

	cmds = exec.Command(initSymbol, "init") // 子进程的启动命令：1.执行进程内的可执行文件，2.初始化
	cmds.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}

	if isStd {
		cmds.Stdin = os.Stdin
		cmds.Stdout = os.Stdout
		cmds.Stderr = os.Stderr
	} else {
		// 创建日志文件
		dir := fmt.Sprintf(container.GhnDockerRunningContainerDir, containerId)
		if mkdirErr := os.MkdirAll(dir, 0622); mkdirErr != nil {
			logrus.Errorf("mk container log dir failed, err:%s", err)
			return nil, nil
		}
		docUrl := dir + "/" + container.LogFileName

		file, createErr := os.Create(docUrl)
		if createErr != nil {
			logrus.Errorf("create log file failed, err:%s", err)
			return nil, nil
		}

		cmds.Stdout = file
	}

	cmds.ExtraFiles = []*os.File{read}
	cmds.Env = append(env, os.Environ()...)
	cmds.Dir = "/mnt/" + containerId
	if err := container.NewWorkSpace(image, containerId, volume); err != nil {
		return nil, nil
	}

	return cmds, write

}

func Run(isStd bool, cmds []string, conf *subsystem.SubSystemConfig, image string, volume string, name string, env []string, net string, portMapping string) {

	// id
	containerId := randStringBytes(10)

	// 父进程执行内容
	parent, writePipe := fork(isStd, image, containerId, volume, env)
	if err := parent.Start(); err != nil {
		logrus.Fatalf("fork start failed err:%s", err)
	}

	// 持久化单host上的container信息
	containerInfo, recordErr := recordContainerInfo(containerId, image, name, strconv.Itoa(parent.Process.Pid), cmds, volume, net, portMapping)
	if recordErr != nil {
		logrus.Fatalf("record container failed, err:%s", recordErr)
	}

	// 资源限制
	containManager := cgroup.NewCgroupManager(fmt.Sprintf(container.CGroupPathFormat, containerId), conf)
	defer containManager.Remove()
	containManager.ProcessId = strconv.Itoa(parent.Process.Pid)

	err := containManager.ApplySubsystem()
	if err != nil {
		logrus.Fatalf("[containManager.ApplySubsystem] err failed, err:%s", err)
	}

	err = containManager.SetPidIntoGroup()
	if err != nil {
		logrus.Fatalf("[containManager.SetPidIntoGroup] err failed, err:%s", err)
	}

	// 联入指定网络
	if err = network.Connect(net, portMapping, containerInfo); err != nil {
		logrus.Fatal("[network.Connect] container connect network failed, err:%s", err)
	}
	// 执行指令通过管道
	sendInitCommand(cmds, writePipe)

	//原来parent.Wait（）主要是用于父进程等待子进程结束，这在交互式创建容器的步骤里面是没问题的，
	//但是在这里，如果detach创建了容器，就不能再去等待，创建容器之后，父进程就已经退出了。
	// 因此，这里只是将容器内的init进程启动起来，就已经完成工作，紧接着就可以退出，然后由操作系统进程ID为1的init进程去接管容器进程。
	if isStd {
		parent.Wait()
	}
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	logrus.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}

func randStringBytes(n int) string {
	letterBytes := "1234567890"
	rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func recordContainerInfo(containerId, image, name, pid string, cmds []string, volume string, net string, portMapping string) (*container.ContainerInfo, error) {
	if name == "" {
		name = containerId
	}

	containerInfo := &container.ContainerInfo{
		Id:            containerId,
		ContainerName: name,
		Image:         image,
		Pid:           pid,
		Commands:      strings.Join(cmds, " "),
		Status:        container.ContainerStatus_Running,
		CreateTime:    time.Now().Format("2006-01-02 15:04:05"),
		Volume:        volume,
		Network:       net,
		PortMapping:   portMapping,
	}

	// 序列化
	str, err := sonic.Marshal(containerInfo)
	if err != nil {
		logrus.Errorf("[recordContainerInfo] json marshal failed, err:%s", err)
		return nil, err
	}

	docUrl := fmt.Sprintf(container.GhnDockerRunningContainerDir, containerId)
	if err = os.MkdirAll(docUrl, 0777); err != nil {
		logrus.Errorf("[recordContainerInfo] mkdir failed, err:%s", err)
		return nil, err
	}

	docUrl = docUrl + "/" + container.ConfFileName

	file, err := os.Create(docUrl)
	defer file.Close()
	if err != nil {
		logrus.Errorf("[recordContainerInfo] create new file failed, err:%s", err)
		return nil, err
	}

	if _, err = file.WriteString(string(str)); err != nil {
		logrus.Errorf("[recordContainerInfo] write container record info failed, err:%s", err)
		return nil, err
	}

	return containerInfo, nil
}

package container

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/common-tools-haonan/docker/cgroup"
	"github.com/common-tools-haonan/docker/cgroup/subsystem"
	"github.com/common-tools-haonan/docker/network"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"text/tabwriter"
	"time"
)

const (
	CGroupPathFormat             = "/home/guohaonan/ghndocker/container/%s/cgroup"
	GhnDockerRunningContainerDir = "/home/guohaonan/ghndocker/run/%s"
	ConfFileName                 = "config.json"
	LogFileName                  = "container.log"
)

type ContainerStatus string

const (
	ContainerStatus_Running ContainerStatus = "RUNNING"
	ContainerStatus_Stop    ContainerStatus = "STOP"
	ContainerStatus_Exit    ContainerStatus = "EXIT"
)

type ContainerInfo struct {
	Id            string          `json:"id"`
	ContainerName string          `json:"container_name"`
	Pid           string          `json:"pid"`
	Image         string          `json:"image"`
	Status        ContainerStatus `json:"status"`
	Commands      string          `json:"commands"`
	CreateTime    string          `json:"create_time"`
	Volume        string          `json:"volume"`
	Network       string          `json:"network"`
	PortMapping   string          `json:"port_mapping"`
}

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
		dir := fmt.Sprintf(GhnDockerRunningContainerDir, containerId)
		if mkdirErr := os.MkdirAll(dir, 0622); mkdirErr != nil {
			logrus.Errorf("mk container log dir failed, err:%s", err)
			return nil, nil
		}
		docUrl := dir + "/" + LogFileName

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
	if err := NewWorkSpace(image, containerId, volume); err != nil {
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
	recordErr := recordContainerInfo(containerId, image, name, strconv.Itoa(parent.Process.Pid), cmds, volume, net, portMapping)
	if recordErr != nil {
		logrus.Fatalf("record container failed, err:%s", recordErr)
	}

	// 资源限制
	containManager := cgroup.NewCgroupManager(fmt.Sprintf(CGroupPathFormat, containerId), conf)
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
	if err = network.Connect(net, portMapping, containerId); err != nil {
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

func recordContainerInfo(containerId, image, name, pid string, cmds []string, volume string, net string, portMapping string) error {
	if name == "" {
		name = containerId
	}
	containerInfo := &ContainerInfo{
		Id:            containerId,
		ContainerName: name,
		Image:         image,
		Pid:           pid,
		Commands:      strings.Join(cmds, " "),
		Status:        ContainerStatus_Running,
		CreateTime:    time.Now().Format("2006-01-02 15:04:05"),
		Volume:        volume,
		Network:       net,
		PortMapping:   portMapping,
	}

	// 序列化
	str, err := sonic.Marshal(containerInfo)
	if err != nil {
		logrus.Errorf("[recordContainerInfo] json marshal failed, err:%s", err)
		return err
	}

	docUrl := fmt.Sprintf(GhnDockerRunningContainerDir, containerId)
	if err = os.MkdirAll(docUrl, 0777); err != nil {
		logrus.Errorf("[recordContainerInfo] mkdir failed, err:%s", err)
		return err
	}

	docUrl = docUrl + "/" + ConfFileName

	file, err := os.Create(docUrl)
	defer file.Close()
	if err != nil {
		logrus.Errorf("[recordContainerInfo] create new file failed, err:%s", err)
		return err
	}

	if _, err = file.WriteString(string(str)); err != nil {
		logrus.Errorf("[recordContainerInfo] write container record info failed, err:%s", err)
		return err
	}

	return nil
}

func deleteContainerInfo(path string) error {
	if err := os.RemoveAll(path); err != nil {
		logrus.Errorf("delete container Info failed, err:%s", err)
		return err
	}
	return nil
}

// ListAllContainers 将当前host内所有的容器信息输出到标准输出流
func ListAllContainers() {

	files, err := ioutil.ReadDir("/home/guohaonan/ghndocker/run/")
	if err != nil {
		logrus.Errorf("[ListAllContainers] Read Dir failed, err:%s", err)
		return
	}

	containers := make([]*ContainerInfo, 0)
	for _, file := range files {

		var (
			container *ContainerInfo
			tmpErr    error
		)

		container, tmpErr = handleContainerDir(file.Name())
		if tmpErr != nil {
			logrus.Errorf("container:%s handle failed, err:%s", file.Name(), tmpErr)
			continue
		}
		containers = append(containers, container)
	}

	w := tabwriter.NewWriter(os.Stdout, 12, 1, 3, ' ', 0)
	fmt.Fprint(w, "ID\tNAME\tIMAGE\tPID\tSTATUS\tCREATE_TIME\tCMDS\n")

	for _, container := range containers {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
			container.Id,
			container.ContainerName,
			container.Image,
			container.Pid,
			container.Status,
			container.CreateTime,
			container.Commands)
	}

	if err := w.Flush(); err != nil {
		logrus.Fatalf("[ListAllContainers] Flush to buffered zone failed, err:%s", err)
	}
	return
}

func handleContainerDir(dir string) (*ContainerInfo, error) {
	path := fmt.Sprintf(GhnDockerRunningContainerDir, dir) + "/" + ConfFileName
	record, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	container := &ContainerInfo{}
	if err = sonic.Unmarshal(record, container); err != nil {
		return nil, err
	}
	return container, nil

}

// FindContainerLog 根据容器id寻找对应的日志文件，并输出到标准输出流
func FindContainerLog(containerId string) {
	path := fmt.Sprintf(GhnDockerRunningContainerDir, containerId) + "/" + LogFileName
	logFile, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.Fatalf("read log from file failed, err:%s", err)
	}
	fmt.Fprint(os.Stdout, string(logFile))
}

// StopContainer 根据容器id kill对应的进程，并修改持久化存储文件
func StopContainer(containerId string) error {
	path := fmt.Sprintf(GhnDockerRunningContainerDir, containerId) + "/" + ConfFileName
	recordFile, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.Errorf("[StopContainer] read record file failed, err:%s", err)
		return err
	}

	container := &ContainerInfo{}
	if err = sonic.Unmarshal(recordFile, container); err != nil {
		return err
	}

	pid, _ := strconv.Atoi(container.Pid)
	// 杀死容器进程
	if err = syscall.Kill(pid, syscall.SIGTERM); err != nil {
		logrus.Errorf("[StopContainer] kill proc failed, err:%s", err)
		return err
	}

	// 更新容器记录
	container.Status = ContainerStatus_Stop
	container.Pid = " "

	bytes, err := sonic.Marshal(container)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(path, bytes, 0622); err != nil {
		return err
	}
	return nil
}

func RemoveContainer(containerId string, isForce bool) error {
	path := fmt.Sprintf(GhnDockerRunningContainerDir, containerId)
	recordUrl := fmt.Sprintf(GhnDockerRunningContainerDir, containerId) + "/" + ConfFileName
	recordFile, err := ioutil.ReadFile(recordUrl)
	if err != nil {
		logrus.Errorf("[RemoveContainer] read record file failed, err:%s", err)
		return err
	}

	container := &ContainerInfo{}
	if err = sonic.Unmarshal(recordFile, container); err != nil {
		return err
	}

	if !isForce && container.Status != ContainerStatus_Stop {
		logrus.Infof("[RemoveContainer] unforcibly remove only apply for container which has been stop")
		return nil
	}

	// 卸除挂载点
	if err = RemoveMountVolume(containerId); err != nil {
		logrus.Infof("[RemoveContainer] RemoveMountVolume failed, err:%s", err)
		return err
	}

	// 移除容器记录+容器日志
	if err = deleteContainerInfo(path); err != nil {
		return err
	}

	// 删除容器可写层以及work层
	if err = deleteContainerInfo(fmt.Sprintf(GhnDockerContainerDir, containerId)); err != nil {
		return err
	}

	if err = deleteContainerInfo(fmt.Sprintf(GhnDockerWorkDir, containerId)); err != nil {
		return err
	}

	if err = RemoveMountPoints(containerId); err != nil {
		logrus.Infof("[RemoveContainer] RemoveMountPoints failed, err:%s", err)
		return err
	}
	return nil
}

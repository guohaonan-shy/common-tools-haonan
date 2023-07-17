package container

import (
	"fmt"
	"github.com/common-tools-haonan/docker/cgroup"
	"github.com/common-tools-haonan/docker/cgroup/subsystem"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var (
	CGroupPathFormat = "/home/guohaonan/ghndocker/container/%s/cgroup"
)

func fork(isStd bool, image, containerId, volume string) (cmds *exec.Cmd, write *os.File) {

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
	}

	cmds.ExtraFiles = []*os.File{read}
	cmds.Env = os.Environ()
	cmds.Dir = "/mnt/" + containerId
	if err := NewWorkSpace(image, containerId, volume); err != nil {
		return nil, nil
	}

	return cmds, write

}

func Run(isStd bool, cmds []string, conf *subsystem.SubSystemConfig, image string, volume string) {

	// id
	containerId := randStringBytes(10)

	// 父进程执行内容
	parent, writePipe := fork(isStd, image, containerId, volume)
	if err := parent.Start(); err != nil {
		logrus.Fatalf("fork start failed err:%s", err)
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

	// 执行指令通过管道
	sendInitCommand(cmds, writePipe)

	if isStd {
		parent.Wait()
		os.Remove(path.Join("./" + containerId))
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

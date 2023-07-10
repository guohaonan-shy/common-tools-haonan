package container

import (
	"github.com/common-tools-haonan/docker/cgroup"
	"github.com/common-tools-haonan/docker/cgroup/subsystem"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func fork(isStd bool) (cmds *exec.Cmd, write *os.File) {

	cmds = exec.Command("/proc/self/exe", "init") // 子进程的启动命令：1.执行进程内的可执行文件，2.初始化
	cmds.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}

	if isStd {
		cmds.Stdin = os.Stdin
		cmds.Stdout = os.Stdout
		cmds.Stderr = os.Stderr
	}
	read, write, err := os.Pipe()
	if err != nil {
		logrus.Fatal("the process of creating a pipe failed occurring fork, err:%s ", err)
	}

	cmds.ExtraFiles = []*os.File{read}

	return cmds, write

}

func RunContainer(isStd bool, cmds []string, conf *subsystem.SubSystemConfig) {
	// 父进程执行内容
	parent, writePipe := fork(isStd)
	if err := parent.Start(); err != nil {
		logrus.Error(err)
	}

	// id
	containerId := randStringBytes(18)

	// 资源限制
	containManager := cgroup.NewCgroupManager(containerId, conf)
	containManager.ProcessId = strconv.Itoa(parent.Process.Pid)

	err := containManager.ApplySubsystem()
	if err != nil {
		logrus.Fatal("[containManager.ApplySubsystem] err failed, err:%s", err)
	}

	err = containManager.SetPidIntoGroup()
	if err != nil {
		logrus.Fatal("[containManager.SetPidIntoGroup] err failed, err:%s", err)
	}

	defer containManager.Remove()

	// 执行指令通过管道
	command := strings.Join(cmds, " ")
	logrus.Infof("[write pipe]command all is %s", command)
	if _, err := writePipe.WriteString(command); err != nil {
		logrus.Fatal(err)
	}

	parent.Wait()
	os.Exit(-1)
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

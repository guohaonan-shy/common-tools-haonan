package container

import (
	"github.com/common-tools-haonan/docker/cgroup"
	"github.com/common-tools-haonan/docker/cgroup/subsystem"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"
)

func fork(isStd bool, cmd string) (cmds *exec.Cmd, write *os.File) {

	cmds = exec.Command("/proc/self/exe", "init") // 子进程的启动命令：1.执行进程内的可执行文件，2.初始化
	cmds.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}

	if isStd {
		cmds.Stdin = os.Stdin
		cmds.Stdout = os.Stdout
		cmds.Stderr = os.Stderr
	}
	write, read, err := os.Pipe()
	if err != nil {
		logrus.Fatal("the process of creating a pipe failed occurring fork, err:%s ", err)
	}

	cmds.ExtraFiles = []*os.File{read}

	return cmds, write

}

func RunContainer(isStd bool, cmd string, conf *subsystem.SubSystemConfig) {
	// 父进程执行内容
	cmds, writePipe := fork(isStd, cmd)
	if err := cmds.Start(); err != nil {
		logrus.Error(err)
	}

	// id
	containerId := randStringBytes(18)

	// 资源限制
	containManager := cgroup.NewCgroupManager(containerId, conf)
	containManager.ProcessId = strconv.Itoa(cmds.Process.Pid)

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
	if _, err := writePipe.WriteString(cmd); err != nil {
		logrus.Fatal(err)
	}

	cmds.Wait()
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

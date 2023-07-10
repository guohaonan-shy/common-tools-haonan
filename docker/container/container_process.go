package container

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"syscall"
)

func Fork(isStd bool) (cmds *exec.Cmd, write *os.File) {

	read, write, err := os.Pipe()
	if err != nil {
		logrus.Fatal("the process of creating a pipe failed occurring fork, err:%s ", err)
	}

	initCmds, err := os.Readlink("/proc/self/exe")
	if err != nil {
		logrus.Fatal(err)
	}
	cmds = exec.Command(initCmds, "init") // 子进程的启动命令：1.执行进程内的可执行文件，2.初始化
	cmds.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}

	if isStd {
		cmds.Stdin = os.Stdin
		cmds.Stdout = os.Stdout
		cmds.Stderr = os.Stderr
	}

	cmds.ExtraFiles = []*os.File{read}

	return cmds, write

}

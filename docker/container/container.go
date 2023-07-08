package container

import (
	"github.com/common-tools-haonan/docker/container/subsystem"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
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

func RunContainer(isStd bool, cmd string) {
	// 父进程执行内容
	cmds, writePipe := fork(isStd, cmd)
	if err := cmds.Start(); err != nil {
		logrus.Error(err)
	}

	// 执行指令通过管道
	if _, err := writePipe.WriteString(cmd); err != nil {
		logrus.Fatal(err)
	}

	cmds.Wait()
	os.Exit(-1)
}

// 创建子进程是否，命令行输入是/proc/self/exe init;即先执行父进程的所有可执行内容，然后执行init
// 这块我觉得有点巧合，在于github.com/urfave/cli这个包本身在run的过程中，就将系统输入的第一个参数自动默认为cli的name，所以这个init本质上是ghndocker init的实现
func RunContainerInitProcess() error {

	var (
		msg []byte
	)

	readPipe := os.NewFile(uintptr(3), "pipe")
	_, err := readPipe.Read(msg)
	if err != nil {
		logrus.Errorf("read from read pipe failed, err:%s", err)
		os.Exit(-1)
	}
	logrus.Infof("command %s", string(msg))

	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")

	// 寻找命令行工具的可执行文件
	cmdArrays := strings.Split(string(msg), " ")
	path, err := exec.LookPath(cmdArrays[0])
	if err != nil {
		logrus.Errorf("find cmd binary path failed, err:%s", err)
		os.Exit(-1)
	}

	if err := syscall.Exec(path, cmdArrays[1:], os.Environ()); err != nil {
		logrus.Errorf(err.Error())
	}
	return nil
}

func RunContainerWithConfig(isStd bool, cmd string, namespace string, conf *subsystem.SubSystemConfig) {
	parent, _ := fork(isStd, cmd)
	if err := parent.Start(); err != nil {
		logrus.Fatal(err)
	}

	defer func() {
		// 释放资源
		for _, subIns := range subsystem.SubSystemFactory {
			subIns.Remove(namespace)
		}
	}()

	for _, subIns := range subsystem.SubSystemFactory {
		// 创建cgroup，设置配置
		err := subIns.Apply(namespace, conf)
		if err != nil {
			logrus.Fatal("apply failed")
		}
		// 设置pid到task
		err = subIns.SetPid(namespace, strconv.Itoa(parent.Process.Pid))
		if err != nil {
			logrus.Fatal("SetPid failed")
		}
	}
	RunContainer(isStd, cmd)

	parent.Wait()
}

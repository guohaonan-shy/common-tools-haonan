package container

import (
	"github.com/common-tools-haonan/docker/container/subsystem"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func fork(isStd bool, cmd string) *exec.Cmd {
	args := []string{"init", cmd}
	osCmd := exec.Command("/proc/self/exe", args...)
	osCmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}

	if isStd {
		osCmd.Stdin = os.Stdin
		osCmd.Stdout = os.Stdout
		osCmd.Stderr = os.Stderr
	}
	return osCmd

}

func RunContainer(isStd bool, cmd string) {
	parentProcess := fork(isStd, cmd)
	if err := parentProcess.Start(); err != nil {
		logrus.Error(err)
	}
	parentProcess.Wait()
	os.Exit(-1)
}
func RunContainerInitProcess(cmd string, args []string) error {
	logrus.Infof("command %s", cmd)

	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	argv := []string{cmd}
	if err := syscall.Exec(cmd, argv, os.Environ()); err != nil {
		logrus.Errorf(err.Error())
	}
	return nil
}

func RunContainerWithConfig(isStd bool, cmd string, namespace string, conf *subsystem.SubSystemConfig) {
	parent := fork(isStd, cmd)
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
		// 设置pid到task
		err = subIns.SetPid(namespace, strconv.Itoa(parent.Process.Pid))
	}
	RunContainer(isStd, cmd)

	parent.Wait()
}

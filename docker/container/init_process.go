package container

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

// 创建子进程是否，命令行输入是/proc/self/exe init;即先执行父进程的所有可执行内容，然后执行init
// 这块我觉得有点巧合，在于github.com/urfave/cli这个包本身在run的过程中，就将系统输入的第一个参数自动默认为cli的name，所以这个init本质上是ghndocker init的实现
func RunContainerInitProcess() error {

	var (
		msg []byte
	)
	// read pipe，无内容阻塞后面处理逻辑l
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

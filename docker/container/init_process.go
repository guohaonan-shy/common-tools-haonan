package container

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
)

// 创建子进程是否，命令行输入是/proc/self/exe init;即先执行父进程的所有可执行内容，然后执行init
// 这块我觉得有点巧合，在于github.com/urfave/cli这个包本身在run的过程中，就将系统输入的第一个参数自动默认为cli的name，所以这个init本质上是ghndocker init的实现
func RunContainerInitProcess() error {

	// read pipe，无内容阻塞后面处理逻辑l
	cmds := readUserCommand()
	if cmds == nil || len(cmds) == 0 {
		return fmt.Errorf("Run container get user command error, cmdArray is nil")
	}
	logrus.Infof("command %s", cmds)

	if err := setupMount(); err != nil {
		os.Exit(-1)
	}

	// 寻找命令行工具的可执行文件
	execPath, err := exec.LookPath(cmds[0])
	if err != nil {
		logrus.Errorf("find cmd binary path failed, err:%s", err)
		os.Exit(-1)
	}

	if err := syscall.Exec(execPath, cmds[0:], os.Environ()); err != nil {
		logrus.Errorf(err.Error())
	}
	return nil
}

func readUserCommand() []string {
	pipe := os.NewFile(uintptr(3), "pipe")
	defer pipe.Close()
	msg, err := io.ReadAll(pipe)
	if err != nil {
		logrus.Errorf("init read pipe error %v", err)
		return nil
	}
	msgStr := string(msg)
	return strings.Split(msgStr, " ")
}

func setupMount() error {
	wd, err := os.Getwd()
	if err != nil {
		logrus.Errorf("[setupMount] get current work directory failed, err:%s", err)
		return fmt.Errorf("when container is initing, get current work directory failed, err:%s", err)
	}

	err = pivotRoot(wd)
	if err != nil {
		return fmt.Errorf("[setupMount] pivot root failed, err:%s", err)
	}

	//mount proc
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")

	syscall.Mount("tmpfs", "/dev", "tmpfs", syscall.MS_NOSUID|syscall.MS_STRICTATIME, "mode=755")

	return nil
}

func pivotRoot(root string) error {
	// 这块解释一下，为啥出现mount -b挂载源和目标相同
	// 原因主要是因为pivot_root这个系统调用，新的root必须是mount挂载点，来自该系统调用的约束
	if err := syscall.Mount(root, root, "bind", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		logrus.Errorf("[pivotRoot] mount bind root failed, err:%s", err)
		return fmt.Errorf("pivot_root failed, err:%s", err)
	}

	// 创建pivot_root系统调用的另一个目录，put_old
	putOld := path.Join(root, ".pivot_root")
	if err := os.Mkdir(putOld, 0777); err != nil {
		logrus.Errorf("[pivotRoot] create temporary directory failed, err:%s", err)
		return fmt.Errorf("pivot_root failed, err:%s", err)
	}

	// 这块有个点在于，pivot_root是保证，调用的进程或线程所在的namespace包含的所有进程或线程的根目录都是new_root，即初始化之后容器生命周期所有该namespace的进程均是new_root作为工作根目录
	if err := syscall.PivotRoot(root, putOld); err != nil {
		logrus.Errorf("[pivotRoot] system call pivot_root failed, err:%s", err)
		return fmt.Errorf("pivot_root failed, err:%s", err)
	}

	// 切换工作的根目录
	if err := os.Chdir("/"); err != nil {
		logrus.Errorf("[pivotRoot] change failed, err:%s", err)
		return fmt.Errorf("pivot_root failed, err:%s", err)
	}

	putOld = path.Join("/", ".pivot_root")
	if err := syscall.Unmount(putOld, syscall.MNT_DETACH); err != nil {
		logrus.Errorf("[pivotRoot] unmount failed, err:%s", err)
		return fmt.Errorf("pivot_root failed, err:%s", err)
	}

	// 移除临时文件夹
	return os.Remove(putOld)
}

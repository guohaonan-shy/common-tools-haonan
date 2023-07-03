package demo

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

// 理解namespace的demo代码
func Namespace() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 1234,
				HostID:      syscall.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 1234,
				HostID:      syscall.Getgid(),
				Size:        1,
			},
		},
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

// 理解cgroup的demo代码
const cgroupMemoryHierarchyMount = "/sys/fs/cgroup/memory"

func CGroup() {
	// proc/self 类比于 proc/$$ 获取当前进程的信息
	if os.Args[0] == "/proc/self/exe" {
		fmt.Printf("current pid:%d", syscall.Getpid())

		cmd := exec.Command("sh", "-c", "stress --vm-bytes 200m --vm-keep -m 1")
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	cmd := exec.Command("proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		// 得到fork出来的进程映射在外部命名空间的pid
		fmt.Printf("%v", cmd.Process.Pid)

		//
		os.Mkdir(path.Join(cgroupMemoryHierarchyMount, "test_memory_limit"), 0755)

		//
		os.WriteFile(path.Join(cgroupMemoryHierarchyMount, "test_memory_limit", "tasks"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)

		//
		os.WriteFile(path.Join(cgroupMemoryHierarchyMount, "test_memory_limit", "memory.limit_in_bytes"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
	}
	cmd.Process.Wait()

}

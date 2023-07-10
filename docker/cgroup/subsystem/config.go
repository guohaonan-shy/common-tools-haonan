package subsystem

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

var (
	SubSystemFactory = []ContainerSubsystem{
		&MemoryConfig{},
		&CpuShareConfig{},
		&CpuSetConfig{},
	}
)

type ContainerSubsystem interface {
	// 返回对应子系统的名称
	Name() string
	// 将创建的容器pid放入指定的
	SetPid(namespace string, pid string) error
	// 将对应的子系统配置加载更新到对应的cgroup内
	Apply(namespace string, conf *SubSystemConfig) error
	// 移除子系统内部某个对应的cgroup
	Remove(namespace string) error
}

type SubSystemConfig struct {
	MemoryLimits string
	CpuShare     string
	CpuSet       string
}

func findRootPathBySubsystem(subsystem string) (string, error) {
	file, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return "", fmt.Errorf("[FindRootPathBySubsystem] io open files failed, err:%s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		txt := scanner.Text()
		fields := strings.Split(txt, " ")

		for _, str := range strings.Split(fields[len(fields)-1], ",") {
			if subsystem == str {
				return fields[4], nil
			}
		}
	}

	// 不会出现EOF
	if scanner.Err() != nil {
		return "", fmt.Errorf("[FindRootPathBySubsystem] the process of scanning files by line occurs err:%s", err)
	}

	return "", nil
}

func getCgroupPathWithCreateOption(subsystem string, namespace string, autoCreate bool) (string, error) {
	root, _ := findRootPathBySubsystem(subsystem)
	if _, err := os.Stat(path.Join(root, namespace)); err == nil || (autoCreate && os.IsNotExist(err)) {
		if os.IsNotExist(err) {
			if err := os.Mkdir(path.Join(root, namespace), 0755); err == nil {
			} else {
				return "", fmt.Errorf("error create cgroup %v", err)
			}
		}
		return path.Join(root, namespace), nil
	} else {
		return "", fmt.Errorf("cgroup path error %v", err)
	}
}

package subsystem

import (
	"fmt"
	"os"
	"path"
)

const (
	SubsystemName_CpuNumber = "CpuNumber"
)

type CpuNumberConfig struct {
}

func NewCpuNumberConfig() *CpuNumberConfig {
	return &CpuNumberConfig{}
}

func (cpuNumber *CpuNumberConfig) Name() string {
	return SubsystemName_CpuNumber
}

func (cpuNumber *CpuNumberConfig) SetPid(namespace string, pid string) error {
	root, err := getCgroupPathWithCreateOption(cpuNumber.Name(), namespace, false)
	if err != nil {
		return err
	}
	// 将进程pid放入group底下的tasks文件
	err = os.WriteFile(path.Join(root, "tasks"), []byte(pid), 0644)
	if err != nil {
		return fmt.Errorf("[CpuNumberConfig] write files failed, err:%s", err)
	}
	return nil
}

func (cpuNumber *CpuNumberConfig) Apply(namespace string, conf *SubSystemConfig) error {
	root, err := getCgroupPathWithCreateOption(cpuNumber.Name(), namespace, true)
	if err != nil {
		return err
	}
	// 将内存配置放入group底下的cpuset.cpus文件
	err = os.WriteFile(path.Join(root, "cpuset.cpus"), []byte(conf.CpuNumber), 0644)
	if err != nil {
		return fmt.Errorf("[CpuNumberConfig] write files failed, err:%s", err)
	}
	return nil
}

func (cpuNumber *CpuNumberConfig) Remove(namespace string) error {
	root, err := getCgroupPathWithCreateOption(cpuNumber.Name(), namespace, true)
	if err != nil {
		return err
	}
	return os.Remove(root)
}

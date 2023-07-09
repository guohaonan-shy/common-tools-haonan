package subsystem

import (
	"fmt"
	"os"
	"path"
)

const (
	SubsystemName_CpuSet = "CpuSet"
)

type CpuSetConfig struct {
}

func NewCpuSetConfig() *CpuSetConfig {
	return &CpuSetConfig{}
}

func (cpuSet *CpuSetConfig) Name() string {
	return SubsystemName_CpuSet
}

func (cpuSet *CpuSetConfig) SetPid(namespace string, pid string) error {
	root, err := getCgroupPathWithCreateOption(cpuSet.Name(), namespace, false)
	if err != nil {
		return err
	}
	// 将进程pid放入group底下的tasks文件
	err = os.WriteFile(path.Join(root, "tasks"), []byte(pid), 0644)
	if err != nil {
		return fmt.Errorf("[CpuSetConfig] write files failed, err:%s", err)
	}
	return nil
}

func (cpuSet *CpuSetConfig) Apply(namespace string, conf *SubSystemConfig) error {
	root, err := getCgroupPathWithCreateOption(cpuSet.Name(), namespace, true)
	if err != nil {
		return err
	}
	// 将内存配置放入group底下的cpuset.cpus文件
	err = os.WriteFile(path.Join(root, "cpu.shares"), []byte(conf.CpuSet), 0644)
	if err != nil {
		return fmt.Errorf("[CpuSetConfig] write files failed, err:%s", err)
	}
	return nil
}

func (cpuSet *CpuSetConfig) Remove(namespace string) error {
	root, err := getCgroupPathWithCreateOption(cpuSet.Name(), namespace, false)
	if err != nil {
		return err
	}
	return os.Remove(root)
}

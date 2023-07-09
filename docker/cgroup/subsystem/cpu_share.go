package subsystem

import (
	"fmt"
	"os"
	"path"
)

const (
	SubsystemName_CpuShare = "CpuShare"
)

type CpuShareConfig struct {
}

func NewCpuShareConfig() *CpuShareConfig {
	return &CpuShareConfig{}
}

func (cpuShare *CpuShareConfig) Name() string {
	return SubsystemName_CpuShare
}

func (cpuShare *CpuShareConfig) SetPid(namespace string, pid string) error {
	root, err := getCgroupPathWithCreateOption(cpuShare.Name(), namespace, false)
	if err != nil {
		return err
	}
	// 将进程pid放入group底下的tasks文件
	err = os.WriteFile(path.Join(root, "tasks"), []byte(pid), 0644)
	if err != nil {
		return fmt.Errorf("[CpuShareConfig] write files failed, err:%s", err)
	}
	return nil
}

func (cpuShare *CpuShareConfig) Apply(namespace string, conf *SubSystemConfig) error {
	root, err := getCgroupPathWithCreateOption(cpuShare.Name(), namespace, true)
	if err != nil {
		return err
	}
	// 将内存配置放入group底下的cpuset.cpus文件
	err = os.WriteFile(path.Join(root, "cpuset.cpus"), []byte(conf.CpuShare), 0644)
	if err != nil {
		return fmt.Errorf("[CpuShareConfig] write files failed, err:%s", err)
	}
	return nil
}

func (cpuShare *CpuShareConfig) Remove(namespace string) error {
	root, err := getCgroupPathWithCreateOption(cpuShare.Name(), namespace, false)
	if err != nil {
		return err
	}
	return os.Remove(root)
}

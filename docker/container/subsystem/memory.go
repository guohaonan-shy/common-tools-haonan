package subsystem

import (
	"fmt"
	"os"
	"path"
)

const (
	SubsystemName_Memory = "Memory"
)

type MemoryConfig struct {
}

func NewMemoryConfig() *MemoryConfig {
	return &MemoryConfig{}
}

func (memory *MemoryConfig) Name() string {
	return SubsystemName_Memory
}

func (memory *MemoryConfig) SetPid(namespace string, pid string) error {
	root, err := getCgroupPathWithCreateOption(memory.Name(), namespace, false)
	if err != nil {
		return err
	}
	// 将进程pid放入group底下的tasks文件
	err = os.WriteFile(path.Join(root, "tasks"), []byte(pid), 0644)
	if err != nil {
		return fmt.Errorf("[MemoryConfig] write files failed, err:%s", err)
	}
	return nil
}

func (memory *MemoryConfig) Apply(namespace string, conf *SubSystemConfig) error {
	root, err := getCgroupPathWithCreateOption(memory.Name(), namespace, true)
	if err != nil {
		return err
	}
	// 将内存配置放入group底下的memory.limit_in_bytes文件
	err = os.WriteFile(path.Join(root, "memory.limit_in_bytes"), []byte(conf.MemoryLimits), 0644)
	if err != nil {
		return fmt.Errorf("[MemoryConfig] write files failed, err:%s", err)
	}
	return nil
}

func (memory *MemoryConfig) Remove(namespace string) error {
	root, err := getCgroupPathWithCreateOption(memory.Name(), namespace, true)
	if err != nil {
		return err
	}
	return os.Remove(root)
}

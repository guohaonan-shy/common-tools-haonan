package cgroup

import (
	"github.com/common-tools-haonan/docker/cgroup/subsystem"
	"github.com/sirupsen/logrus"
)

type CgroupManager struct {
	// Namespace 代表相对于/sys/fd/cgroup/{具体资源}/{namespace}，即hierarchy的树状层级关系
	Namespace string
	// SubSystemConf 进程cgroup的子系统以及子系统配置
	SubSystemConf *subsystem.SubSystemConfig
	// ProcessId 新fork出来的进程id
	ProcessId string
}

func NewCgroupManager(namespace string, conf *subsystem.SubSystemConfig) *CgroupManager {
	return &CgroupManager{
		Namespace:     namespace,
		SubSystemConf: conf,
	}
}

func (manager *CgroupManager) ApplySubsystem() error {

	var (
		namespace     = manager.Namespace
		subsystemConf = manager.SubSystemConf
	)

	for _, subIns := range subsystem.SubSystemFactory {
		// 创建cgroup，设置配置
		err := subIns.Apply(namespace, subsystemConf)
		if err != nil {
			logrus.Errorf("apply failed")
			return err
		}
	}
	return nil
}

func (manager *CgroupManager) SetPidIntoGroup() error {
	var (
		namespace = manager.Namespace
	)

	for _, subIns := range subsystem.SubSystemFactory {
		// 设置pid到task
		err := subIns.SetPid(namespace, manager.ProcessId)
		if err != nil {
			logrus.Errorf("SetPid failed")
			return err
		}
	}
	return nil
}

func (manager *CgroupManager) Remove() error {
	var (
		namespace = manager.Namespace
	)
	// 释放资源
	for _, subIns := range subsystem.SubSystemFactory {
		err := subIns.Remove(namespace)
		if err != nil {
			return err
		}
	}
	return nil
}

//func RunContainerWithConfig(isStd bool, cmd string, namespace string, conf *subsystem.SubSystemConfig) {
//	parent, _ := fork(isStd, cmd)
//	if err := parent.Start(); err != nil {
//		logrus.Fatal(err)
//	}
//
//	defer func() {
//		// 释放资源
//		for _, subIns := range subsystem.SubSystemFactory {
//			subIns.Remove(namespace)
//		}
//	}()
//
//	for _, subIns := range subsystem.SubSystemFactory {
//		// 创建cgroup，设置配置
//		err := subIns.Apply(namespace, conf)
//		if err != nil {
//			logrus.Fatal("apply failed")
//		}
//		// 设置pid到task
//		err = subIns.SetPid(namespace, strconv.Itoa(parent.Process.Pid))
//		if err != nil {
//			logrus.Fatal("SetPid failed")
//		}
//	}
//	RunContainer(isStd, cmd)
//
//	parent.Wait()
//}

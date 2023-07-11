package main

import (
	"github.com/common-tools-haonan/docker/cgroup"
	"github.com/common-tools-haonan/docker/cgroup/subsystem"
	"github.com/common-tools-haonan/docker/container"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func Run(isStd bool, cmds []string, conf *subsystem.SubSystemConfig) {
	// 父进程执行内容
	parent, writePipe := container.Fork(isStd)
	if err := parent.Start(); err != nil {
		logrus.Error(err)
	}

	// id
	containerId := randStringBytes(10)

	// 资源限制
	containManager := cgroup.NewCgroupManager(containerId, conf)
	defer containManager.Remove()
	containManager.ProcessId = strconv.Itoa(parent.Process.Pid)

	err := containManager.ApplySubsystem()
	if err != nil {
		logrus.Fatal("[containManager.ApplySubsystem] err failed, err:%s", err)
	}

	err = containManager.SetPidIntoGroup()
	if err != nil {
		logrus.Fatal("[containManager.SetPidIntoGroup] err failed, err:%s", err)
	}

	// 执行指令通过管道
	sendInitCommand(cmds, writePipe)

	if isStd {
		parent.Wait()
		os.Remove(path.Join("./" + containerId))
	}
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	logrus.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}

func randStringBytes(n int) string {
	letterBytes := "1234567890"
	rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

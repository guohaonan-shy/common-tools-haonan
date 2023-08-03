package container

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
	"text/tabwriter"
)

const (
	CGroupPathFormat             = "/home/guohaonan/ghndocker/container/%s/cgroup"
	GhnDockerRunningContainerDir = "/home/guohaonan/ghndocker/run/%s"
	ConfFileName                 = "config.json"
	LogFileName                  = "container.log"
)

type ContainerStatus string

const (
	ContainerStatus_Running ContainerStatus = "RUNNING"
	ContainerStatus_Stop    ContainerStatus = "STOP"
	ContainerStatus_Exit    ContainerStatus = "EXIT"
)

type ContainerInfo struct {
	Id            string          `json:"id"`
	ContainerName string          `json:"container_name"`
	Pid           string          `json:"pid"`
	Image         string          `json:"image"`
	Status        ContainerStatus `json:"status"`
	Commands      string          `json:"commands"`
	CreateTime    string          `json:"create_time"`
	Volume        string          `json:"volume"`
	Network       string          `json:"network"`
	PortMapping   string          `json:"port_mapping"`
}

func deleteContainerInfo(path string) error {
	if err := os.RemoveAll(path); err != nil {
		logrus.Errorf("delete container Info failed, err:%s", err)
		return err
	}
	return nil
}

// ListAllContainers 将当前host内所有的容器信息输出到标准输出流
func ListAllContainers() {

	files, err := ioutil.ReadDir("/home/guohaonan/ghndocker/run/")
	if err != nil {
		logrus.Errorf("[ListAllContainers] Read Dir failed, err:%s", err)
		return
	}

	containers := make([]*ContainerInfo, 0)
	for _, file := range files {

		var (
			container *ContainerInfo
			tmpErr    error
		)

		container, tmpErr = handleContainerDir(file.Name())
		if tmpErr != nil {
			logrus.Errorf("container:%s handle failed, err:%s", file.Name(), tmpErr)
			continue
		}
		containers = append(containers, container)
	}

	w := tabwriter.NewWriter(os.Stdout, 12, 1, 3, ' ', 0)
	fmt.Fprint(w, "ID\tNAME\tIMAGE\tPID\tSTATUS\tCREATE_TIME\tCMDS\n")

	for _, container := range containers {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
			container.Id,
			container.ContainerName,
			container.Image,
			container.Pid,
			container.Status,
			container.CreateTime,
			container.Commands)
	}

	if err := w.Flush(); err != nil {
		logrus.Fatalf("[ListAllContainers] Flush to buffered zone failed, err:%s", err)
	}
	return
}

func handleContainerDir(dir string) (*ContainerInfo, error) {
	path := fmt.Sprintf(GhnDockerRunningContainerDir, dir) + "/" + ConfFileName
	record, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	container := &ContainerInfo{}
	if err = sonic.Unmarshal(record, container); err != nil {
		return nil, err
	}
	return container, nil

}

// FindContainerLog 根据容器id寻找对应的日志文件，并输出到标准输出流
func FindContainerLog(containerId string) {
	path := fmt.Sprintf(GhnDockerRunningContainerDir, containerId) + "/" + LogFileName
	logFile, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.Fatalf("read log from file failed, err:%s", err)
	}
	fmt.Fprint(os.Stdout, string(logFile))
}

// StopContainer 根据容器id kill对应的进程，并修改持久化存储文件
func StopContainer(containerId string) error {
	path := fmt.Sprintf(GhnDockerRunningContainerDir, containerId) + "/" + ConfFileName
	recordFile, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.Errorf("[StopContainer] read record file failed, err:%s", err)
		return err
	}

	container := &ContainerInfo{}
	if err = sonic.Unmarshal(recordFile, container); err != nil {
		return err
	}

	pid, _ := strconv.Atoi(container.Pid)
	// 杀死容器进程
	if err = syscall.Kill(pid, syscall.SIGTERM); err != nil {
		logrus.Errorf("[StopContainer] kill proc failed, err:%s", err)
		return err
	}

	// 更新容器记录
	container.Status = ContainerStatus_Stop
	container.Pid = " "

	bytes, err := sonic.Marshal(container)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(path, bytes, 0622); err != nil {
		return err
	}
	return nil
}

func RemoveContainer(containerId string, isForce bool) error {
	path := fmt.Sprintf(GhnDockerRunningContainerDir, containerId)
	recordUrl := fmt.Sprintf(GhnDockerRunningContainerDir, containerId) + "/" + ConfFileName
	recordFile, err := ioutil.ReadFile(recordUrl)
	if err != nil {
		logrus.Errorf("[RemoveContainer] read record file failed, err:%s", err)
		return err
	}

	container := &ContainerInfo{}
	if err = sonic.Unmarshal(recordFile, container); err != nil {
		return err
	}

	if !isForce && container.Status != ContainerStatus_Stop {
		logrus.Infof("[RemoveContainer] unforcibly remove only apply for container which has been stop")
		return nil
	}

	// 卸除挂载点
	if err = RemoveMountVolume(containerId); err != nil {
		logrus.Infof("[RemoveContainer] RemoveMountVolume failed, err:%s", err)
		return err
	}

	// 移除容器记录+容器日志
	if err = deleteContainerInfo(path); err != nil {
		return err
	}

	// 删除容器可写层以及work层
	if err = deleteContainerInfo(fmt.Sprintf(GhnDockerContainerDir, containerId)); err != nil {
		return err
	}

	if err = deleteContainerInfo(fmt.Sprintf(GhnDockerWorkDir, containerId)); err != nil {
		return err
	}

	if err = RemoveMountPoints(containerId); err != nil {
		logrus.Infof("[RemoveContainer] RemoveMountPoints failed, err:%s", err)
		return err
	}
	return nil
}

func GetSpecificContainers(containerId string) (*ContainerInfo, error) {

	containerFile, err := os.ReadFile("/home/guohaonan/ghndocker/run/" + containerId)
	if err != nil {
		logrus.Errorf("[getSpecificContainers] Read file failed, err:%s", err)
		return nil, err
	}

	container := &ContainerInfo{}
	if err = sonic.Unmarshal(containerFile, container); err != nil {
		return nil, err
	}
	return container, nil
}

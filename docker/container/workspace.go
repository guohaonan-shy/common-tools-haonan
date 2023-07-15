package container

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

const (
	FileSystem_OverlayFormat = "lowerdir=%s,upperdir=%s,workdir=%s %s"

	GhnDockerImageDir     = "/home/guohaonan/ghndocker/image/%s"
	GhnDockerContainerDir = "/home/guohaonan/ghndocker/container/%s"
	GhnDockerWorkDir      = "/home/guohaonan/ghndocker/work/%s"
	GhnDockerMountPoint   = "/mnt/%s"
)

func NewWorkSpace(image string, containerId string) error {

	var (
		imageErr, containerErr, mntErr error
	)

	defer func() {
		if containerErr != nil {
			RemoveContainerLayer(containerId)
			return
		}

		if mntErr != nil {
			RemoveContainerLayer(containerId)
			RemoveMountPoints(containerId)
			return
		}
		return
	}()

	if imageErr = CreateImageLayer(image); imageErr != nil {
		return imageErr
	}

	if containerErr = CreateContainerLayer(containerId); containerErr != nil {
		return containerErr
	}

	if mntErr = CreateMountPoints(image, containerId); mntErr != nil {
		return mntErr
	}

	return nil
}

func CreateImageLayer(image string) error {
	imageUrl := fmt.Sprintf(GhnDockerImageDir, image)

	isExist, err := PathExist(imageUrl)
	if err != nil {
		logrus.Errorf("[CreateImageLayer] check image url existence failed, err:%s", err)
		return err
	}

	if isExist {
		logrus.Info("image has existed")
		return nil
	}

	// 镜像文件不存在
	if err = os.Mkdir(imageUrl, 0777); err != nil {
		logrus.Errorf("[CreateImageLayer] mk image dir failed, err:%s", err)
		return err
	}

	imageTarUrl := fmt.Sprintf(GhnDockerImageDir, image) + ".tar"
	if _, err := exec.Command("tar", "xvf", imageTarUrl, "-C", imageUrl).CombinedOutput(); err != nil {
		logrus.Errorf("tar cmd exec failed, err:%s", err)
		return err
	}
	return nil
}

func CreateContainerLayer(containerId string) error {
	containerUrl := fmt.Sprintf(GhnDockerContainerDir, containerId)
	if err := os.MkdirAll(containerUrl, 0777); err != nil {
		logrus.Errorf("[CreateContainerLayer] mk container dir failed, err:%s", err)
		return err
	}

	tmpWorkUrl := fmt.Sprintf(GhnDockerWorkDir, containerId)
	if err := os.MkdirAll(tmpWorkUrl, 0777); err != nil {
		logrus.Errorf("[CreateContainerLayer] mk tmp work dir failed, err:%s", err)
		return err
	}
	return nil
}

func CreateMountPoints(image string, containerId string) error {
	mountUrl := fmt.Sprintf(GhnDockerMountPoint, containerId)
	if err := os.MkdirAll(mountUrl, 0777); err != nil {
		logrus.Errorf("[CreateMountPoints] mk mnt dir failed, err:%s", err)
		return err
	}

	imageUrl, containerUrl := fmt.Sprintf(GhnDockerImageDir, image), fmt.Sprintf(GhnDockerContainerDir, containerId)
	tmpWorkUrl := fmt.Sprintf(GhnDockerWorkDir, containerId)

	dirs := fmt.Sprintf(FileSystem_OverlayFormat, imageUrl, containerUrl, tmpWorkUrl, mountUrl)

	logrus.Infof("aufs dirs:%s", dirs)
	// mount -t overlay -o lowerdir=./lower,upperdir=./upper,workdir=./work ./merged
	if out, err := exec.Command("mount", "-t", "overlay", "overlay", "-o", dirs).CombinedOutput(); err != nil {
		logrus.Errorf("mount failed, err:%s \n stdout:%s", err, string(out))
		return err
	}

	return nil
}

func PathExist(url string) (bool, error) {
	_, err := os.Stat(url)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func RemoveContainerLayer(containerId string) error {
	containerUrl := fmt.Sprintf(GhnDockerContainerDir, containerId)
	if err := os.RemoveAll(containerUrl); err != nil {
		logrus.Errorf("[RemoveContainerLayer] rm container dir failed, err:%s", err)
		return err
	}

	tmpWorkUrl := fmt.Sprintf(GhnDockerWorkDir, containerId)
	if err := os.RemoveAll(tmpWorkUrl); err != nil {
		logrus.Errorf("[RemoveContainerLayer] mk tmp work dir failed, err:%s", err)
		return err
	}
	return nil
}

func RemoveMountPoints(containerId string) error {
	mountUrl := fmt.Sprintf(GhnDockerMountPoint, containerId)
	if err := os.RemoveAll(mountUrl); err != nil {
		logrus.Errorf("[RemoveMountPoints] rm mnt dir failed, err:%s", err)
		return err
	}

	return nil
}

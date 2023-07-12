package container

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

func NewWorkSpace(image string, containerId string) error {
	if err := CreateImageLayer(image); err != nil {
		return err
	}

	if err := CreateContainerLayer(containerId); err != nil {
		return err
	}

	if err := CreateMountPoints(image, containerId); err != nil {
		return err
	}

	return nil
}

func CreateImageLayer(image string) error {
	imageUrl := "/home/guohaonan/image/" + image

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

	imageTarUrl := "/home/guohaonan/image/" + image + ".tar"
	if _, err := exec.Command("tar", "xvf", imageTarUrl, "-C", imageUrl).CombinedOutput(); err != nil {
		logrus.Errorf("tar cmd exec failed, err:%s", err)
		return err
	}
	return nil
}

func CreateContainerLayer(containerId string) error {
	containerUrl := "home/guohaonan/container/" + containerId
	if err := os.MkdirAll(containerUrl, 0777); err != nil {
		logrus.Errorf("[CreateContainerLayer] mk container dir failed, err:%s", err)
		return err
	}
	return nil
}

func CreateMountPoints(image string, containerId string) error {
	mountUrlFormat := "/mnt/%s/"
	mountUrl := fmt.Sprintf(mountUrlFormat, containerId)
	if err := os.MkdirAll(mountUrl, 0777); err != nil {
		logrus.Errorf("[CreateMountPoints] mk mnt dir failed, err:%s", err)
		return err
	}

	imageUrl := "/home/guohaonan/image/" + image
	containerUrl := "/home/guohaonan/container/" + containerId

	dirs := "dirs=" + containerUrl + ":" + imageUrl

	logrus.Infof("aufs dirs:%s", dirs)
	if out, err := exec.Command("mount", "-t", "overlay", "-o", dirs, "none", mountUrl).CombinedOutput(); err != nil {
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

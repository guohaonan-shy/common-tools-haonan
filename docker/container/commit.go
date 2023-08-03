package container

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

func CommitToMakeAImage(containerId string, image string) error {
	containerFileUrl := fmt.Sprintf(GhnDockerMountPoint, containerId)
	imageName := image + ".tar"
	imageUrl := fmt.Sprintf(GhnDockerImageDir, imageName)
	_, err := os.Stat(imageUrl)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if output, err := exec.Command("tar", "-czf", imageUrl, "-C", containerFileUrl, ".").CombinedOutput(); err != nil {
		logrus.Errorf("commit image failed, err:%s, output:%s", err, output)
		return err
	}
	return nil
}

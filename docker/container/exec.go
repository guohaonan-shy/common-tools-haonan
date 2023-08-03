package container

/*
#define _GNU_SOURCE
#include <unistd.h>
#include <errno.h>
#include <sched.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>

__attribute__((constructor)) void enter_namespace(void) {
	char *ghndocker_pid;
	ghndocker_pid = getenv("ghndocker_pid");
	if (ghndocker_pid) {
		fprintf(stdout, "got ghndocker_pid=%s\n", ghndocker_pid);
	} else {
		fprintf(stdout, "missing ghndocker_pid env skip nsenter");
		return;
	}
	char *ghndocker_cmd;
	ghndocker_cmd = getenv("ghndocker_cmd");
	if (ghndocker_cmd) {
		fprintf(stdout, "got ghndocker_cmd=%s\n", ghndocker_cmd);
	} else {
		fprintf(stdout, "missing ghndocker_cmd env skip nsenter");
		return;
	}
	int i;
	char nspath[1024];
	char *namespaces[] = { "ipc", "uts", "net", "pid", "mnt" };

	for (i=0; i<5; i++) {
		sprintf(nspath, "/proc/%s/ns/%s", ghndocker_pid, namespaces[i]);
		int fd = open(nspath, O_RDONLY);

		if (setns(fd, 0) == -1) {
			fprintf(stderr, "setns on %s namespace failed: %s\n", namespaces[i], strerror(errno));
		} else {
			fprintf(stdout, "setns on %s namespace succeeded\n", namespaces[i]);
		}
		close(fd);
	}
	int res = system(ghndocker_cmd);
	exit(0);
	return;
}
*/
import "C"
import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const (
	ENV_EXEC_PID = "ghndocker_pid"
	ENV_EXEC_CMD = "ghndocker_cmd"
)

func ExecContainer(containerId string, cmds []string) error {
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

	pid := container.Pid
	cmd := strings.Join(cmds, " ")

	os.Setenv(ENV_EXEC_CMD, cmd)
	os.Setenv(ENV_EXEC_PID, pid)

	execCmd := exec.Command("/proc/self/exe", "exec")
	execCmd.Env = append(os.Environ(), getEnvsByPid(pid)...)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	execCmd.Stdin = os.Stdin

	if err = execCmd.Run(); err != nil {
		logrus.Errorf("/proc/self/exe -exec failed, err:%s", err)
		return err
	}
	return nil
}

func getEnvsByPid(pid string) []string {
	path := fmt.Sprintf("/proc/%s/environ", pid)
	contentBytes, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.Errorf("Read file %s error %v", path, err)
		return nil
	}
	//env split by \u0000
	envs := strings.Split(string(contentBytes), "\u0000")
	return envs
}

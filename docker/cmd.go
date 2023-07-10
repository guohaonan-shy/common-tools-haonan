package main

import (
	"fmt"
	"github.com/common-tools-haonan/docker/cgroup/subsystem"
	"github.com/common-tools-haonan/docker/container"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "ghndocker"
	app.Usage = "ghndocker is a simple docker cmdline tool for guohaonan.Aatrox use"
	app.Commands = []cli.Command{
		initCommand,
		runCommand,
	}

	app.Before = func(context *cli.Context) error {
		// log
		logrus.SetFormatter(&logrus.JSONFormatter{})

		logrus.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init cgroup process run user's process in cgroup. Don't call it outside",
	Action: func(context *cli.Context) error {
		logrus.Infof("init start")
		return container.RunContainerInitProcess()
	},
}

var runCommand = cli.Command{
	Name: "run",
	Usage: "Create a cgroup with namespace and cgroups limit " +
		"ghndocker run -it [command] ",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "it",
			Usage: "enable docker to run",
		},
		cli.StringFlag{
			Name:  "memory",
			Usage: "memory limit",
		},
		cli.StringFlag{
			Name:  "cpushre",
			Usage: "cpushare limit",
		},
		cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpuset limit",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("missing Container command")
		}
		cmd := context.Args().Get(0)
		itFlag := context.Bool("it")

		resConf := &subsystem.SubSystemConfig{
			MemoryLimits: context.String("memory"),
			CpuSet:       context.String("cpuset"),
			CpuShare:     context.String("cpushare"),
		}

		container.RunContainer(itFlag, cmd, resConf)
		return nil
	},
}

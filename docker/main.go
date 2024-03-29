package main

import (
	"fmt"
	"github.com/common-tools-haonan/docker/cgroup/subsystem"
	"github.com/common-tools-haonan/docker/container"
	"github.com/common-tools-haonan/docker/network"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"text/tabwriter"
)

func main() {

	//
	network.Init()

	app := cli.NewApp()
	app.Name = "ghndocker"
	app.Usage = "ghndocker is a simple docker cmdline tool for guohaonan.Aatrox use"
	app.Commands = []cli.Command{
		initCommand,
		runCommand,
		listCommand,
		logCommand,
		stopCommand,
		removeCommand,
		execCommand,
		commitCommand,
		networkCommand,
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
		cli.BoolFlag{
			Name:  "detach",
			Usage: "run container on background",
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
		cli.StringFlag{
			Name:  "image",
			Usage: "image type and name",
		},
		cli.StringFlag{
			Name:  "volume",
			Usage: "mount volume between host and container",
		},
		cli.StringFlag{
			Name:  "name",
			Usage: "container name",
		},
		cli.StringSliceFlag{
			Name:  "env",
			Usage: "environment from stdin",
		},
		cli.StringFlag{
			Name:  "net",
			Usage: "connect into the specific network",
		},
		cli.StringFlag{
			Name:  "port",
			Usage: "host port mapping with container port",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("missing Container command")
		}
		cmds := make([]string, 0)

		for _, cmd := range context.Args() {
			cmds = append(cmds, cmd)
		}

		itFlag := context.Bool("it")
		detachFlag := context.Bool("detach")

		if itFlag && detachFlag {
			logrus.Fatalf("itFlag and detachFlag cannot exist at the same time")
		}

		resConf := &subsystem.SubSystemConfig{
			MemoryLimits: context.String("memory"),
			CpuSet:       context.String("cpuset"),
			CpuShare:     context.String("cpushare"),
		}

		image := context.String("image")
		volume := context.String("volume")
		name := context.String("name")
		env := context.StringSlice("env")
		net := context.String("net")
		portMapping := context.String("port")

		Run(itFlag, cmds, resConf, image, volume, name, env, net, portMapping)
		return nil
	},
}

var listCommand = cli.Command{
	Name:  "show",
	Usage: "show all containers in this host",
	Action: func(context *cli.Context) error {
		container.ListAllContainers()
		return nil
	},
}

var logCommand = cli.Command{
	Name:  "logs",
	Usage: "show exec logs in container",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "container_id",
			Usage: "find log by container_id",
		},
	},
	Action: func(ctx *cli.Context) {
		id := ctx.String("container_id")
		container.FindContainerLog(id)
		return
	},
}

var stopCommand = cli.Command{
	Name:  "stop",
	Usage: "stop container",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "container_id",
			Usage: "stop container by container_id",
		},
	},
	Action: func(ctx *cli.Context) error {
		id := ctx.String("container_id")
		return container.StopContainer(id)
	},
}

var removeCommand = cli.Command{
	Name:  "remove",
	Usage: "remove container and its' documents",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "container_id",
			Usage: "remove container by container_id",
		},
		cli.BoolFlag{
			Name:  "f",
			Usage: "remove container and its' documents forcibly",
		},
	},
	Action: func(ctx *cli.Context) error {
		force := ctx.Bool("f")
		return container.RemoveContainer(ctx.String("container_id"), force)
	},
}

var execCommand = cli.Command{
	Name:  "exec",
	Usage: "exec a command into container",
	Action: func(context *cli.Context) error {
		//This is for callback
		if os.Getenv(container.ENV_EXEC_PID) != "" {
			logrus.Infof("pid callback pid %s", os.Getgid())
			return nil
		}

		if len(context.Args()) < 2 {
			return fmt.Errorf("Missing container name or command")
		}
		containerName := context.Args().Get(0)
		var commandArray []string
		for _, arg := range context.Args().Tail() {
			commandArray = append(commandArray, arg)
		}
		container.ExecContainer(containerName, commandArray)
		return nil
	},
}

var commitCommand = cli.Command{
	Name:  "commit",
	Usage: "commit a image based on a existed container filesystem",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "container_id",
			Usage: "container filesystem based on",
		},
		cli.StringFlag{
			Name:  "image",
			Usage: "image name",
		},
	},
	Action: func(context *cli.Context) error {
		containerId, image := context.String("container_id"), context.String("image")
		return container.CommitToMakeAImage(containerId, image)
	},
}

var networkCommand = cli.Command{
	Name:  "network",
	Usage: "tools about container network, for example, create network(LAN)",
	Subcommands: []cli.Command{
		{
			Name:  "create",
			Usage: "Create Network based on the specific driver",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "name",
					Usage:    "network name",
					Required: true,
				},
				cli.StringFlag{
					Name:     "driver",
					Usage:    "driver name, network config based on driver",
					Required: true,
				},
				cli.StringFlag{
					Name:     "subnet",
					Usage:    "subnet",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				name, driver, subnet := ctx.String("name"), ctx.String("driver"), ctx.String("subnet")
				return network.CreateNetwork(name, driver, subnet)
			},
		},
		{
			Name:  "list",
			Usage: "list all networks",
			Action: func(ctx *cli.Context) error {
				networks, err := network.ListAllNetwork()
				if err != nil {
					return err
				}

				w := tabwriter.NewWriter(os.Stdout, 12, 1, 3, ' ', 0)
				fmt.Fprintf(w, "NETWORK\tIPRANGE\tDRIVER\n")

				for i := range networks {
					nw := networks[i]
					fmt.Fprintf(w, "%s\t%s\t%s\n", nw.NetworkName, nw.IPRange.String(), nw.Driver)
				}

				return w.Flush()
			},
		},
		{
			Name:  "delete",
			Usage: "delete Network",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "name",
				},
			},
			Action: func(ctx *cli.Context) error {
				name := ctx.String("name")
				return network.DeleteNetwork(name)
			},
		},
	},
}

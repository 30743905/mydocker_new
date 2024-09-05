//go:build linux
// +build linux

package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"mydocker/container"
)

var runCommand = cli.Command{
	Name: "run",
	Usage: `Create a container with namespace and cgroups limit
			mydocker run -it [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "it", // 简单起见，这里把 -i 和 -t 参数合并成一个
			Usage: "enable tty",
		},
	},
	/*
			这里是run命令执行的真正函数。
			1.判断参数是否包含command
			2.获取用户指定的command
			3.调用Run function去准备启动容器:

		main [global options] command [command options] [arguments...]
	*/
	Action: func(context *cli.Context) error {
		log.Info("run args:", context.Args(), "-----", os.Getpid())
		if len(context.Args()) < 1 {
			return fmt.Errorf("missing container command")
		}
		cmd := context.Args().Get(0) //取参数(arguments)
		tty := context.Bool("it")    //取command options
		Run(tty, cmd)
		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	/*
		1.获取传递过来的 command 参数
		2.执行容器初始化操作
	*/
	Action: func(context *cli.Context) error {
		log.Info("init come on", "-----", os.Getpid())
		log.Info("init receive args:", context.Args())
		cmd := context.Args().Get(0)
		log.Infof("command: %s", cmd)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}

package main

import (
	"Odysseus/cmd/init"
	"Odysseus/cmd/run"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

const usage = `Ulysses：is a simple container runtime implementation.
			   The purpose of this project is to learn how docker works and how to write a docker by ourselves
			   Enjoy it, just for fun.`

func main() {
	app := cli.NewApp()
	app.Name = "Ulysses"
	app.Usage = usage

	//命令行处理 共有两个命令
	app.Commands = []cli.Command{
		//启动容器
		run.RunCommand,
		//真正的启动方法
		init.InitCommand,
	}

	app.Before = func(context *cli.Context) error {
		// Log as JSON instead of the default ASCII formatter.
		log.SetFormatter(&log.JSONFormatter{})

		log.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

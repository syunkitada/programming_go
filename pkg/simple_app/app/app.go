package app

import (
	"os"

	"github.com/syunkitada/go-samples/pkg/simple_app/config"
	"github.com/urfave/cli"

	"github.com/golang/glog"
)

var (
	Conf = &config.Conf
)

func Main() error {
	cli.VersionFlag = config.VersionFlag

	app := cli.NewApp()
	app.Name = "sample-simple-app"
	app.Usage = "sample-simple-app"
	app.Version = "0.0.1"
	app.Flags = append(config.CommonFlags, config.GlogFlags...)

	app.Action = func(c *cli.Context) error {
		config.Init(c)
		glog.Info(Conf.Api)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		glog.Fatal(err)
		return err
	}

	return nil
}

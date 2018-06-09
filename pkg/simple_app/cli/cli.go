package cli

import (
	"os"

	"github.com/syunkitada/go-sample/pkg/simple_app/config"
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

	app.Commands = []cli.Command{
		{
			Name:  "status",
			Usage: "usage",
			Action: func(c *cli.Context) error {
				config.Init(c)
				glog.Info("Execute status")
				return nil
			},
		},
		{
			Name:  "test",
			Usage: "usage",
			Action: func(c *cli.Context) error {
				config.Init(c)
				glog.Info("Execute test")
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		glog.Fatal(err)
		return err
	}

	return nil
}

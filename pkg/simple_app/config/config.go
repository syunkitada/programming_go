package config

import (
	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
	"github.com/syunkitada/go-sample/pkg/simple_app/testdata"
	"github.com/urfave/cli"
)

var Conf Config

var CommonFlags = []cli.Flag{
	cli.StringFlag{Name: "config-file", Value: testdata.Path("conf.toml"), Usage: "config-file"},
}

var VersionFlag = cli.BoolFlag{Name: "print-version, V", Usage: "print only the version"}

func Init(ctx *cli.Context) {
	glogGangstaShim(ctx)
	newConfig := newConfig(ctx)
	_, err := toml.DecodeFile(ctx.GlobalString("config-file"), newConfig)
	if err != nil {
		glog.Errorf("Failed to decode file : %!s(MISSING)", err)
		return
	}
	Conf = *newConfig
	glog.Infof("Loaded config-file(%v)", ctx.GlobalString("config-file"))
}

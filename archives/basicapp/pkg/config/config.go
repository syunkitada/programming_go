package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

var (
	configDir string
)

var (
	glogV               int
	glogLogtostderr     bool
	glogStderrthreshold int
	glogAlsologtostderr bool
	glogVmodule         string
	glogLogDir          string
	glogLogBacktraceAt  string
)

var Conf Config

func InitFlags(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().StringVar(&configDir, "config-dir", "", "config directory (default is $HOME/.etc)")
	rootCmd.PersistentFlags().IntVar(&glogV, "glog-v", 0, "log level for V logs")
	rootCmd.PersistentFlags().BoolVar(&glogLogtostderr, "glog-logtostderr", true, "log to standard error instead of files")
	rootCmd.PersistentFlags().IntVar(&glogStderrthreshold, "glog-stderrthreshold", 0, "logs at or above this threshold go to stderr")
	rootCmd.PersistentFlags().BoolVar(&glogAlsologtostderr, "glog-alsologtostderr", false, "log to standard error as well as files")
	rootCmd.PersistentFlags().StringVar(&glogVmodule, "glog-vmodule", "", "comma-separated list of pattern=N settings for file-filtered logging")
	rootCmd.PersistentFlags().StringVar(&glogLogDir, "glog-log-dir", "", "If non-empty, write log files in this directory")
	rootCmd.PersistentFlags().StringVar(&glogLogBacktraceAt, "glog-log-backtrace-at", ":0", "when logging hits line file:N, emit a stack trace")
}

func InitConfig() {
	_ = flag.CommandLine.Parse([]string{})
	flagShim(map[string]string{
		"v":                fmt.Sprint(glogV),
		"logtostderr":      fmt.Sprint(glogLogtostderr),
		"stderrthreshold":  fmt.Sprint(glogStderrthreshold),
		"alsologtostderr":  fmt.Sprint(glogAlsologtostderr),
		"vmodule":          glogVmodule,
		"log_dir":          glogLogDir,
		"log_backtrace_at": glogLogBacktraceAt,
	})

	if configDir == "" {
		pwd := os.Getenv("PWD")
		configDir = filepath.Join(pwd, "ci", "etc")
	}

	err := loadConfig(configDir)
	if err != nil {
		glog.Fatal(err)
	}
}

func flagShim(fakeVals map[string]string) {
	flag.VisitAll(func(fl *flag.Flag) {
		if val, ok := fakeVals[fl.Name]; ok {
			fl.Value.Set(val)
		}
	})
}

func loadConfig(configDir string) error {
	newConfig := newConfig(configDir)
	configFile := filepath.Join(configDir, "app.toml")
	_, err := toml.DecodeFile(configFile, newConfig)
	if err != nil {
		return err
	}
	Conf = *newConfig

	return nil
}

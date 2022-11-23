package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	// MEMO: github.com/ghodss/yaml の Unmarshalは、デフォルト値を引き継げる
	"github.com/ghodss/yaml"
)

func MustLoadConfigFiles(conf interface{}, defaultConf interface{}, configFiles []string) {
	if err := LoadConfigFiles(conf, defaultConf, configFiles); err != nil {
		fmt.Fprintf(os.Stderr, "Failed MustLoadConfigFiles: %s", strings.Join(configFiles, ","))
		os.Exit(1)
	}
}

func LoadConfigFiles(conf interface{}, defaultConf interface{}, configFiles []string) (err error) {
	var bytes []byte
	if bytes, err = yaml.Marshal(defaultConf); err != nil {
		return
	}
	if err = yaml.Unmarshal(bytes, &conf); err != nil {
		return
	}

	for _, configFile := range configFiles {
		if bytes, err = ioutil.ReadFile(configFile); err != nil {
			return
		}
		if err = yaml.Unmarshal(bytes, &conf); err != nil {
			return
		}
	}
	return
}

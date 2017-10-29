package conf

import (
	"errors"
	"os"
	"strings"

	"github.com/go-ini/ini"
)

var (
	Conf              config
	defaultConfigFile = "conf/rum.conf"
)

type config struct {
	EnabledSuffix []string
	Sphinx        []string
	Logfile       string
}

func InitConfig(configFile string) error {
	if configFile == "" {
		configFile = defaultConfigFile
	}
	Conf = config{}

	if _, err := os.Stat(configFile); err != nil {
		return errors.New("config file err:" + err.Error())
	} else {
		cfg := ini.Empty()
		err := cfg.Append(defaultConfigFile)
		if err != nil {
			return errors.New("config load err:" + err.Error())
		}

		enabledSuffix := cfg.Section("default").Key("enabled_suffix").MustString("")
		sphinx := cfg.Section("default").Key("sphinx").MustString("")
		logfile := cfg.Section("default").Key("logfile").MustString("/var/log/rum.log")
		Conf.EnabledSuffix = strings.Split(enabledSuffix, ",")
		Conf.Sphinx = strings.Split(sphinx, ",")
		Conf.Logfile = logfile
	}
	return nil
}

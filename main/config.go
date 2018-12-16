package main

import (
	"errors"
	"fmt"

	"logagent/tailf"

	"github.com/astaxie/beego/config"
)

var (
	appConfig *Config
)

type Config struct {
	logLevel string
	logPath  string

	chanSize  int
	kafkaAddr string

	collectConf []tailf.CollectConf
}

func loadConf(confType, filename string) error {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Printf("new config failed, err: %v\n", err)
		return err
	}

	appConfig = &Config{}
	appConfig.logLevel = conf.String("logs::log_level")
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "debug"
	}

	appConfig.logPath = conf.String("logs::log_path")
	if len(appConfig.logPath) == 0 {
		appConfig.logPath = "./logs"
	}

	appConfig.chanSize, err = conf.Int("collect::chan_size")
	if err != nil {
		appConfig.chanSize = 100
	}

	err = loadCollectConf(conf)
	if err != nil {
		fmt.Printf("load collect conf failed, err: %v\n", err)
		return err
	}
	return nil
}

func loadCollectConf(conf config.Configer) error {
	var cc CollectConf
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err := errors.New("invaild collect:log_path")
		return err
	}

	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		err := errors.New("invaild collect:topic")
		return err
	}

	appConfig.collectConf = append(appConfig.collectConf, cc)
	return nil
}

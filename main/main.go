package main

import (
	"fmt"
	"logagent/kafka"
	"logagent/tailf"

	"github.com/astaxie/beego/logs"
	"github.com/prometheus/common/log"
)

func main() {
	filename := "./config/logagent.conf"
	err := loadConf("ini", filename)
	if err != nil {
		fmt.Printf("load conf failed, err: %v\n", err)
		panic("load conf failed")
		return
	}

	err = initLogger()
	if err != nil {
		fmt.Printf("load logger failed, err: %v\n", err)
		panic("load logger failed")
		return
	}
	logs.Debug("load conf success, config: %v", appConfig)

	err = tailf.InitTail(appConfig.collectConf)
	if err != nil {
		logs.Error("init tail failed, err: %v\n", err)
		return
	}

	err = kafka.InitKafka(appConfig.kafkaAddr)
	if err != nil {
		logs.Error("init kafka failed, err: %v\n", err)
		return
	}
	logs.Debug("initialize all success.")

	err = serverRun()
	if err != nil {
		logs.Error("serverRun failed, err: %v", err)
		return
	}

	log.Info("program exited.")
}

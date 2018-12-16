package main

import (
	"logagent/kafka"
	"logagent/tailf"
	"time"

	"github.com/astaxie/beego/logs"
)

func serverRun() error {
	for {
		msg := tailf.GetOneLine()
		err := sendToKafka(msg)
		if err != nil {
			logs.Error("send to kafka failed, err: %v\n", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return nil
}

func sendToKafka(msg *tailf.TextMsg) error {
	//logs.Debug("read msg: %s, topic: %s", msg.Msg, msg.Topic)
	kafka.SendToKafka(msg.Msg, msg.Topic)
	return nil
}

package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	client sarama.SyncProducer
)

func InitKafka(addr string) error {
	config := sarama.NewConfig()
	config.Producer.RequireAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Seccesses = true

	var err error
	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		logs.Error("init kafka producer failed, err:", err)
		return err
	}

	logs.Debug("init kafka success")
	return nil
}

func SendToKafka(data, topic string) error {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringMessage(data)

	_, _, err := client.SendMessage(msg)
	if err != nil {
		logs.Error("send message failed, err: %v, data: %v, topic: %v\n", err, data, topic)
		return err
	}

	//log.Debug("send success, pid: %v, offset: %v, topic: %v\n", pid, offset, topic)
	return nil
}

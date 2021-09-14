package reader

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"

	"github.com/bygui86/go-kafka-segmentio/reader/commons"
	"github.com/bygui86/go-kafka-segmentio/reader/logging"
	"github.com/bygui86/go-kafka-segmentio/reader/monitoring"
)

func (c *KafkaReader) startReader() {
	for {
		msg, err := c.reader.ReadMessage(context.Background())

		monitoring.IncreaseOpsCounter(commons.ServiceName)

		if err == nil { // SUCCESS
			topicInfo, msgInfo := c.getMessageInfo(&msg)
			logging.SugaredLog.Infof("Message received: topic[%s], msg[%s]",
				topicInfo, msgInfo)

			monitoring.IncreaseSuccConsumedMsgCounter(commons.ServiceName, msg.Topic)

		} else { // FAIL
			logging.SugaredLog.Errorf("Failed consuming message:: %s", err.Error())
			monitoring.IncreaseFailConsumedMsgCounter(commons.ServiceName, c.cfg.kafkaTopic)
		}
	}
}

func (c *KafkaReader) getMessageInfo(msg *kafka.Message) (string, string) {
	topicInfo := fmt.Sprintf("name[%s], partition[%d], offset[%d]",
		msg.Topic, msg.Partition, msg.Offset)
	msgInfo := fmt.Sprintf("key[%s], value[%s], timestamp[%v]",
		string(msg.Key), string(msg.Value), msg.Time)
	return topicInfo, msgInfo
}

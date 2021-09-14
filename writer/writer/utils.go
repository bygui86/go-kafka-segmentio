package writer

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"

	"github.com/bygui86/go-kafka-segmentio/writer/commons"
	"github.com/bygui86/go-kafka-segmentio/writer/logging"
	"github.com/bygui86/go-kafka-segmentio/writer/monitoring"
)

// Produce messages to topic (asynchronously)
func (p *KafkaWriter) startWriter() {
	p.ticker = time.NewTicker(1 * time.Second)
	counter := 0

	for {
		select {
		case <-p.ticker.C:
			logging.Log.Info("Tik-tok time to publish")
			// go p.publishMessage(counter)
			p.publishMessage(counter)
		case <-p.stop:
			return
		}
	}
}

func (p *KafkaWriter) publishMessage(counter int) {
	msg := p.messages[counter]

	kmsg := kafka.Message{
		Key:   []byte(fmt.Sprintf("Key-%d", counter)),
		Value: []byte(msg),
	}
	err := p.writer.WriteMessages(context.Background(), kmsg)
	if err != nil {
		logging.SugaredLog.Errorf("Writer failed to publish message %s: %s", msg, err.Error())
		monitoring.IncreaseFailPublishedMsgCounter(commons.ServiceName, p.cfg.kafkaTopic)
		return
	}

	monitoring.IncreaseSuccPublishedMsgCounter(commons.ServiceName, p.cfg.kafkaTopic)
	logging.SugaredLog.Infof("Writer published message %s", msg)

	if counter == len(p.messages)-1 {
		counter = 0
	} else {
		counter++
	}

	monitoring.IncreaseOpsCounter(commons.ServiceName)
}

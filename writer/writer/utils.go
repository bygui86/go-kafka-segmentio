package writer

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"

	"github.com/bygui86/go-kafka-segmentio/writer/commons"
	"github.com/bygui86/go-kafka-segmentio/writer/logging"
	"github.com/bygui86/go-kafka-segmentio/writer/monitoring"
)

// VERSION 1 - OK

func newKafkaWriter(kafkaBroker, kafkaTopic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
}

func startWriterV1(kafkaBroker, kafkaTopic string) {
	writer := newKafkaWriter(kafkaBroker, kafkaTopic)
	defer writer.Close()

	fmt.Println("start producing ... !!")

	for counter := 0; ; counter++ {
		key := fmt.Sprintf("Key-%d", counter)
		msg := kafka.Message{
			Key:   []byte(key),
			Value: []byte(fmt.Sprint(uuid.New())),
		}

		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("produced", key)
		}

		time.Sleep(1 * time.Second)
	}
}

// VERSION 2 - OK

func (p *KafkaWriter) startWriterV2() {
	writer := newKafkaWriter(p.cfg.kafkaBrokers[0], p.cfg.kafkaTopic)
	defer writer.Close()

	logging.SugaredLog.Infof("Start publishing events to %s, topic %s",
		strings.Join(p.cfg.kafkaBrokers, ","), p.cfg.kafkaTopic)

	for counter := 0; ; counter++ {
		key := fmt.Sprintf("Key-%d", counter)
		msg := kafka.Message{
			Key:   []byte(key),
			Value: []byte(fmt.Sprint(uuid.New())),
		}

		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			logging.SugaredLog.Errorf("Failed to publish message '%s': %s", key, err.Error())
		} else {
			logging.SugaredLog.Infof("Published message '%s'", key)
		}

		time.Sleep(1 * time.Second)
	}
}

// VERSION 3 - OK

func (p *KafkaWriter) startWriterV3() {
	logging.SugaredLog.Infof("Start publishing events to %s, topic %s",
		strings.Join(p.cfg.kafkaBrokers, ","), p.cfg.kafkaTopic)

	for counter := 0; ; counter++ {
		key := fmt.Sprintf("Key-%d", counter)
		msg := kafka.Message{
			Key:   []byte(key),
			Value: []byte(fmt.Sprint(uuid.New())),
		}

		err := p.writer.WriteMessages(context.Background(), msg)
		if err != nil {
			logging.SugaredLog.Errorf("Failed to publish message '%s': %s", key, err.Error())
		} else {
			logging.SugaredLog.Infof("Published message '%s'", key)
		}

		time.Sleep(1 * time.Second)
	}
}

// VERSION 4 - OK

func (p *KafkaWriter) startWriterV4() {
	logging.SugaredLog.Infof("Start publishing events to %s, topic %s",
		strings.Join(p.cfg.kafkaBrokers, ","), p.cfg.kafkaTopic)

	for counter := 0; ; {
		logging.Log.Info("Tik-tok time to publish")

		p.publishMessage(counter)

		if counter == len(p.messages)-1 {
			counter = 0
		} else {
			counter++
		}

		time.Sleep(1 * time.Second)
	}
}

// VERSION 5 - OK

func (p *KafkaWriter) startWriterV5() {
	logging.SugaredLog.Infof("Start publishing events to %s, topic %s",
		strings.Join(p.cfg.kafkaBrokers, ","), p.cfg.kafkaTopic)

	for counter := 0; ; {
		logging.Log.Info("Tik-tok time to publish")

		go p.publishMessage(counter)

		if counter == len(p.messages)-1 {
			counter = 0
		} else {
			counter++
		}

		time.Sleep(1 * time.Second)
	}
}

// VERSION FINAL - TODO

// Produce messages to topic
func (p *KafkaWriter) startWriterFinal() {
	logging.SugaredLog.Infof("Start publishing events to %s, topic %s",
		strings.Join(p.cfg.kafkaBrokers, ","), p.cfg.kafkaTopic)

	p.ticker = time.NewTicker(1 * time.Second)
	counter := 0

	for {
		select {
		case <-p.ticker.C:
			// logging.Log.Info("Tik-tok time to publish")

			go p.publishMessage(counter)

			if counter == len(p.messages)-1 {
				counter = 0
			} else {
				counter++
			}

		case <-p.stop:
			return
		}
	}
}

// COMMONS

func (p *KafkaWriter) publishMessage(counter int) {
	text := p.messages[counter]

	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("Key-%d", counter)),
		Value: []byte(text),
	}
	err := p.writer.WriteMessages(context.Background(), msg)
	if err != nil {
		logging.SugaredLog.Errorf("Writer failed to publish message '%s': %s", text, err.Error())
		monitoring.IncreaseFailPublishedMsgCounter(commons.ServiceName, p.cfg.kafkaTopic)
		return
	}

	monitoring.IncreaseSuccPublishedMsgCounter(commons.ServiceName, p.cfg.kafkaTopic)
	logging.SugaredLog.Infof("Writer published message '%s'", text)

	monitoring.IncreaseOpsCounter(commons.ServiceName)
}

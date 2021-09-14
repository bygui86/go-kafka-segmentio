package reader

import (
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"

	"github.com/bygui86/go-kafka-segmentio/reader/logging"
)

func New(consumerName string) (*KafkaReader, error) {
	logging.Log.Info("Create new kafka reader")

	cfg := loadConfig()

	reader := kafka.NewReader(
		kafka.ReaderConfig{
			Brokers: cfg.kafkaBrokers,
			GroupID: cfg.kafkaConsumerGroup,
			Topic:   cfg.kafkaTopic,
			// MinBytes: 10e3, // 10KB
			// MaxBytes: 10e6, // 10MB
		},
	)

	return &KafkaReader{
		cfg:     cfg,
		name:    consumerName,
		reader:  reader,
		running: false,
	}, nil
}

func (c *KafkaReader) Start() error {
	logging.Log.Info("Start kafka reader")

	if c.reader != nil {
		go c.startReader()
		c.running = true
		logging.Log.Info("Kafka reader started")
		return nil
	}

	return fmt.Errorf("kafka reader start failed: reader not initialized or already running")
}

func (c *KafkaReader) Shutdown(timeout int) {
	logging.SugaredLog.Warnf("Shutdown kafka reader, timeout %d", timeout)

	c.stop <- true
	time.Sleep(time.Duration(timeout) * time.Second)
	if c.reader != nil {
		err := c.reader.Close()
		if err != nil {
			logging.SugaredLog.Errorf("Error closing kafka reader: %s", err.Error())
		}
	}
}

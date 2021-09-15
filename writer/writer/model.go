package writer

import (
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaWriter struct {
	cfg      *Config
	name     string
	writer   *kafka.Writer
	ticker   *time.Ticker
	stop     chan bool
	running  bool
	messages []string
}

type Config struct {
	kafkaBrokers []string
	kafkaTopic   string
	kafkaAsync   bool
}

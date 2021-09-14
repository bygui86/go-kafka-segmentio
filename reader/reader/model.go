package reader

import (
	"github.com/segmentio/kafka-go"
)

type KafkaReader struct {
	cfg     *Config
	name    string
	reader  *kafka.Reader
	stop    chan bool
	running bool
}

type Config struct {
	kafkaBrokers       []string
	kafkaTopic         string
	kafkaConsumerGroup string
}

package reader

import (
	"github.com/bygui86/go-kafka-segmentio/reader/logging"
	"github.com/bygui86/go-kafka-segmentio/reader/utils"
)

const (
	kafkaBrokersEnvVar       = "KAFKA_BROKERS" // comma-separated list of 'host:port'
	kafkaTopicEnvVar         = "KAFKA_TOPIC"
	kafkaConsumerGroupEnvVar = "KAFKA_CONSUMER_GROUP"

	kafkaTopicDefault         = "my.topic"
	kafkaConsumerGroupDefault = "my-group"
)

var (
	kafkaBrokersDefault = []string{"localhost:9092"}
)

func loadConfig() *Config {
	logging.Log.Debug("Load kafka producer configurations")
	return &Config{
		kafkaBrokers:       utils.GetStringListEnv(kafkaBrokersEnvVar, kafkaBrokersDefault),
		kafkaTopic:         utils.GetStringEnv(kafkaTopicEnvVar, kafkaTopicDefault),
		kafkaConsumerGroup: utils.GetStringEnv(kafkaConsumerGroupEnvVar, kafkaConsumerGroupDefault),
	}
}

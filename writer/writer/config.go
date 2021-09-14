package writer

import (
	"github.com/bygui86/go-kafka-segmentio/writer/logging"
	"github.com/bygui86/go-kafka-segmentio/writer/utils"
)

const (
	kafkaBrokersEnvVar = "KAFKA_BROKERS" // comma-separated list of 'host:port'
	kafkaTopicEnvVar   = "KAFKA_TOPIC"
	kafkaAsyncEnvVar   = "KAFKA_ASYNC"

	kafkaTopicDefault = "my.topic"
	kafkaAsyncDefault = false
)

var (
	kafkaBrokersDefault = []string{"localhost:9092"}
)

func loadConfig() *Config {
	logging.Log.Debug("Load kafka writer configurations")
	return &Config{
		kafkaBrokers: utils.GetStringListEnv(kafkaBrokersEnvVar, kafkaBrokersDefault),
		kafkaTopic:   utils.GetStringEnv(kafkaTopicEnvVar, kafkaTopicDefault),
		kafkaAsync:   utils.GetBoolEnv(kafkaAsyncEnvVar, kafkaAsyncDefault),
	}
}

package config

import (
	"github.com/bygui86/go-kafka-segmentio/reader/logging"
	"github.com/bygui86/go-kafka-segmentio/reader/utils"
)

const (
	enableMonitoringEnvVar    = "ENABLE_MONITORING"     // bool
	enableCustomMetricsEnvVar = "ENABLE_CUSTOM_METRICS" // bool

	enableMonitoringDefault    = true
	enableCustomMetricsDefault = false
)

func LoadConfig() *Config {
	logging.Log.Info("Load global configurations")

	return &Config{
		enableMonitoring:    utils.GetBoolEnv(enableMonitoringEnvVar, enableMonitoringDefault),
		enableCustomMetrics: utils.GetBoolEnv(enableCustomMetricsEnvVar, enableCustomMetricsDefault),
	}
}

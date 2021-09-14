package config

import (
	"github.com/bygui86/go-kafka-segmentio/writer/logging"
	"github.com/bygui86/go-kafka-segmentio/writer/utils"
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

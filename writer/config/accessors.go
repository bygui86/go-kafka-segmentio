package config

func (c Config) GetEnableMonitoring() bool {
	return c.enableMonitoring
}

func (c Config) GetEnableCustomMetrics() bool {
	return c.enableCustomMetrics
}

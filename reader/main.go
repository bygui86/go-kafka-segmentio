package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bygui86/go-kafka-segmentio/reader/commons"
	"github.com/bygui86/go-kafka-segmentio/reader/config"
	"github.com/bygui86/go-kafka-segmentio/reader/logging"
	"github.com/bygui86/go-kafka-segmentio/reader/monitoring"
	"github.com/bygui86/go-kafka-segmentio/reader/reader"
)

var (
	monitoringServer *monitoring.Server
	kafkaReader      *reader.KafkaReader
)

func main() {
	initLogging()

	logging.SugaredLog.Infof("Start %s", commons.ServiceName)

	cfg := loadConfig()

	if cfg.GetEnableMonitoring() {
		monitoringServer = startMonitoringServer()

		if cfg.GetEnableCustomMetrics() {
			registerCustomMetrics()
		}
	}

	kafkaReader = startReader()

	logging.SugaredLog.Infof("%s up and running", commons.ServiceName)

	startSysCallChannel()

	shutdownAndWait(1)
}

func initLogging() {
	err := logging.InitGlobalLogger()
	if err != nil {
		logging.SugaredLog.Errorf("Logging setup failed: %s", err.Error())
		os.Exit(501)
	}
}

func loadConfig() *config.Config {
	logging.Log.Debug("Load configurations")
	return config.LoadConfig()
}

func startMonitoringServer() *monitoring.Server {
	logging.Log.Debug("Start monitoring")
	server := monitoring.New()
	logging.Log.Debug("Monitoring server successfully created")

	server.Start()
	logging.Log.Debug("Monitoring successfully started")

	return server
}

func registerCustomMetrics() {
	logging.Log.Debug("Register custom metrics")
	monitoring.RegisterCustomMetrics()
}

func startReader() *reader.KafkaReader {
	logging.Log.Debug("Start reader")
	kReader, newErr := reader.New(commons.ServiceName)
	if newErr != nil {
		logging.SugaredLog.Errorf("Consumer setup failed: %s", newErr.Error())
		os.Exit(501)
	}
	logging.Log.Debug("Reader successfully created")

	startErr := kReader.Start()
	if startErr != nil {
		logging.SugaredLog.Errorf("Consumer start failed: %s", startErr.Error())
		os.Exit(502)
	}
	logging.Log.Debug("Reader successfully started")

	return kReader
}

func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
}

func shutdownAndWait(timeout int) {
	logging.SugaredLog.Warnf("Termination signal received! Timeout %d", timeout)

	if kafkaReader != nil {
		kafkaReader.Shutdown(timeout)
	}

	if monitoringServer != nil {
		monitoringServer.Shutdown(timeout)
	}

	time.Sleep(time.Duration(timeout+1) * time.Second)
}

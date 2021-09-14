package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bygui86/go-kafka-segmentio/writer/commons"
	"github.com/bygui86/go-kafka-segmentio/writer/config"
	"github.com/bygui86/go-kafka-segmentio/writer/logging"
	"github.com/bygui86/go-kafka-segmentio/writer/monitoring"
	"github.com/bygui86/go-kafka-segmentio/writer/writer"
)

var (
	monitoringServer *monitoring.Server
	kafkaWriter      *writer.KafkaWriter
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

	kafkaWriter = startWriter()

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

func startWriter() *writer.KafkaWriter {
	logging.Log.Debug("Start writer")
	kWriter, err := writer.New(commons.ServiceName)
	if err != nil {
		logging.SugaredLog.Errorf("Writer setup failed: %s", err.Error())
		os.Exit(501)
	}
	logging.Log.Debug("Writer successfully created")

	kWriter.Start()
	logging.Log.Debug("Writer successfully started")

	return kWriter
}

func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
}

func shutdownAndWait(timeout int) {
	logging.SugaredLog.Warnf("Termination signal received! Timeout %d", timeout)

	if kafkaWriter != nil {
		kafkaWriter.Shutdown(timeout)
	}

	if monitoringServer != nil {
		monitoringServer.Shutdown(timeout)
	}

	time.Sleep(time.Duration(timeout+1) * time.Second)
}

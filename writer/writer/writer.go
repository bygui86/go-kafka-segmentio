package writer

import (
	"time"

	"github.com/segmentio/kafka-go"

	"github.com/bygui86/go-kafka-segmentio/writer/logging"
)

func New(producerName string) (*KafkaWriter, error) {
	logging.Log.Info("Create new kafka writer")

	cfg := loadConfig()

	writer := &kafka.Writer{
		// Addr:  kafka.TCP(cfg.kafkaBrokers...),
		Addr:  kafka.TCP("localhost:9092"),
		Topic: cfg.kafkaTopic,
		// Async: cfg.kafkaAsync,
	}

	return &KafkaWriter{
		cfg:      cfg,
		name:     producerName,
		writer:   writer,
		stop:     make(chan bool, 1),
		running:  false,
		messages: []string{"Frodo Baggins", "Samvise Gamgee", "Meriadoc Brandibuck", "Peregrino Tuc", "Aragorn son of Arathorn", "Boromir son of Denethor", "Legolas son of Thranduil", "Gimli son of Gloin", "Gandalf the Grey"},
	}, nil
}

func (p *KafkaWriter) Start() {
	logging.Log.Info("Start kafka writer")

	if p.writer != nil {
		go p.startWriter()
		p.running = true
		logging.Log.Info("Kafka writer started")
		return
	}

	logging.Log.Error("Kafka writer start failed: writer not initialized or already running")
}

func (p *KafkaWriter) Shutdown(timeout int) {
	logging.SugaredLog.Warnf("Shutdown kafka writer, timeout %d", timeout)

	p.ticker.Stop()
	p.stop <- true
	time.Sleep(time.Duration(timeout) * time.Second)
	if p.writer != nil {
		err := p.writer.Close()
		if err != nil {
			logging.SugaredLog.Errorf("Error closing kafka writer: %s", err.Error())
		}
	}
}

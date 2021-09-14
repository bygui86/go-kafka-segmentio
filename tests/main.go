package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	write()

	time.Sleep(3 * time.Second)

	read()

	// listTopics()

	// time.Sleep(3 * time.Second)

	// produce()

	// time.Sleep(3 * time.Second)

	// consume()
}

func write() {
	// make a writer that produces to topic-A, using the least-bytes distribution
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "my-topic",
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			// Topic: "my-topic",
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			// Topic: "my-topic",
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			// Topic: "my-topic",
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	} else {
		log.Println("Messages written")
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func read() {
	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: "consumer-group-id",
		Topic:   "my-topic",
		// MinBytes:  10e3, // 10KB
		// MaxBytes:  10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func listTopics() {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	log.Println("Connection created")

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	log.Printf("Read %d partitions \n", len(partitions))

	m := map[string]struct{}{}
	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}

	log.Printf("Read %d topics \n", len(m))

	for k := range m {
		fmt.Println(k)
	}
}

// to produce messages
func produce() {
	topic := "my-topic"
	partition := 0

	// conn, connErr := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	conn, connErr := kafka.Dial("tcp", "localhost:9092")
	if connErr != nil {
		log.Fatal("failed to dial leader:", connErr)
	}

	log.Println("Connection created")

	setErr := conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if setErr != nil {
		log.Fatal("failed to set write deadline:", setErr)
	}

	log.Println("Deadline set")

	log.Printf("Producing to topic %s, partition %d \n", topic, partition)

	bytesNum, wrErr := conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if wrErr != nil { // FAIL
		log.Fatal("failed to write messages:", wrErr)
	} else { // SUCCESS
		log.Printf("written %d bytes \n", bytesNum)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

// to consume messages
func consume() {
	topic := "my-topic"
	partition := 0

	conn, connErr := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if connErr != nil {
		log.Fatal("failed to dial leader:", connErr)
	}

	log.Println("Connection created")

	setErr := conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if setErr != nil {
		log.Fatal("failed to set write deadline:", setErr)
	}

	log.Println("Deadline set")

	// batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max
	batch := conn.ReadBatch(1, 1e6)

	log.Println("Batch created")

	log.Printf("Consuming from topic %s, partition %d \n", topic, partition)

	bytesMsg := make([]byte, 1e6) // 10KB max per message
	for {
		_, err := batch.Read(bytesMsg)
		if err != nil {
			log.Printf("ERROR reading message: %+v \n", setErr)
			break
		}
		fmt.Println(string(bytesMsg))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}

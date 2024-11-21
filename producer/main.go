package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	// Kafka broker list
	brokers := []string{"localhost:29092", "localhost:29093", "localhost:29094"}
	topic := "test-topic"

	// Create producer config
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	// Create new producer
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Printf("Error closing producer: %v", err)
		}
	}()

	// Create a context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle OS signals for graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals
		log.Println("Received shutdown signal. Closing producer...")
		cancel()
	}()

	// Produce messages
	counter := 0
	for {
		select {
		case <-ctx.Done():
			return
		default:
			message := fmt.Sprintf("Message %d - %s", counter, time.Now().Format(time.RFC3339))
			partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{
				Topic: topic,
				Value: sarama.StringEncoder(message),
			})

			if err != nil {
				log.Printf("Failed to send message: %v", err)
			} else {
				log.Printf("Message sent successfully! Partition: %d, Offset: %d", partition, offset)
			}

			counter++
			time.Sleep(time.Second)
		}
	}
}

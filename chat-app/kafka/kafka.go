package kafka

import (
    "log"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
    Producer *kafka.Producer // Change to uppercase to export the variable
    topic    = "your_topic_name" // Replace "your_topic_name" with your actual Kafka topic
)

func InitProducer() {
    p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
    if err != nil {
        log.Fatalf("Failed to create Kafka producer: %s\n", err)
    }
    Producer = p // Assign to uppercase variable to export it
}

func ProduceMessage(message []byte) {
    deliveryChan := make(chan kafka.Event)
    defer close(deliveryChan)

    err := Producer.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Value:          message,
    }, deliveryChan)

    if err != nil {
        log.Printf("Failed to produce message to Kafka: %s\n", err)
    }
}

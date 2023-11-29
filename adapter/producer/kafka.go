package producer

import (
	"PostHubApp/domain/entity"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	kafka *kafka.Producer
}

func NewProducer() Producer {
	producer, _ := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9091",
	})
	return Producer{
		kafka: producer,
	}
}

func (producer Producer) Produce(message entity.Migrations) error {
	produceChan := make(chan kafka.Event)
	topic := "POST-HUB-POST-TOPIC"
	messageKafka := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message.ToMessage(),
	}

	err := producer.kafka.Produce(messageKafka, produceChan)

	if err != nil {
		return err
	}

	e := <-produceChan
	if m, ok := e.(*kafka.Message); ok {
		if m.TopicPartition.Error != nil {
			return m.TopicPartition.Error
		}
	} else {
		return fmt.Errorf("unexpected message type: %T", e)
	}

	return nil

}

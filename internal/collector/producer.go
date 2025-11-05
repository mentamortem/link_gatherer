package collector

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewProducer(broker string) (*kafka.Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		return nil, err
	}
	return p, nil

}
func SendLink(producer *kafka.Producer, topic string, message []byte) error {
	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, nil)
	return err
}

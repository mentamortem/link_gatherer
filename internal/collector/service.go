package collector

import (
	"encoding/json"
	"fmt"
	"link_gatherer/pkg/models"
)

func ProccessData(data []byte, format string) error {
	links := models.Link{}
	producer, err := NewProducer("kafka-broker-0.kafka-headless.default.svc.cluster.local:9092")

	if err != nil {
		return fmt.Errorf("Failed to create producer: %w", err)
	}

	switch format {
	case "application/json": // Добавить тест для application/json
		if err := json.Unmarshal(data, &links); err != nil {
			return fmt.Errorf("Failed to unmarshal json: %w", err)
		}
		for _, url := range links.URLS {
			if err := SendLink(producer, "ulrs", []byte(url)); err != nil {
				return fmt.Errorf("Failed to send link to Kafka: %w", err)
			}
		}
	case "text/plain": // Добавить тест для text/plain
		links.URLS = append(links.URLS, string(data))
		for _, url := range links.URLS {
			if err := SendLink(producer, "ulrs", []byte(url)); err != nil {
				return fmt.Errorf("Failed to send link to Kafka: %w", err)
			}
		}
	default:
		return fmt.Errorf("Unsupported format: %s", format)
	}
	return nil
}

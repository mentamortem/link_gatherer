package collector

import (
	"encoding/json"
	"fmt"
	"link_gatherer/pkg/models"
)

func ProccessData(data []byte, format string) error {
	links := models.Link{}
	switch format {
	case "application/json": // Добавить тест для application/json
		if err := json.Unmarshal(data, &links); err != nil {
			return fmt.Errorf("Failed to unmarshal json: %w", err)
		}
	case "text/plain": // Добавить тест для text/plain
		links.URLS = append(links.URLS, string(data))
	default:
		return fmt.Errorf("Unsupported format: %s", format)
	}
	return nil
}

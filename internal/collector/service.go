package collector

import (
	"encoding/json"
	"fmt"
	"link_gatherer/pkg/models"
)

// Если это JSON, то нужно найти в нем поле хранящее ссылки и распарсить их в слайс структуры Link

func ProccessData(data []byte, format string) {
	links := models.Link{}
	switch format {
	case "application/json":
		if err := json.Unmarshal(data, &links); err != nil {
			return
		}
		fmt.Println(links)
	}
}

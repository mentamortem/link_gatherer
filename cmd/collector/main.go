package main

import (
	"fmt"
	"link_gatherer/internal/collector"
)

func main() {
	api := collector.SetupAPI()
	if err := api.Run(":4343"); err != nil {
		fmt.Print("Test")
	}
}

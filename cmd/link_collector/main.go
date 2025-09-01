package main

import (
	"fmt"
	"link_gatherer/internal/collector"
)

func main() {
	collector.Run(8080)
	fmt.Printf("I'm running!")
}

package main

import (
	"fmt"

	"traffic.go/internal/scrape"
)

func main() {

	fmt.Println("Starting TrafficAIO")
	scrape.RunAndSend()

}

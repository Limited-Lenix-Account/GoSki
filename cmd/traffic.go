package main

import (
	"fmt"

	"traffic.go/internal/scrape"
	"traffic.go/util"
)

var MileMarkers []util.MileMarker

func init() {
	MileMarkers = util.ReadMileMarker()
}

func main() {

	// TODO: Start Commenting some stuff to make it more readable in general.
	// 	  Also don't Forget about the travel time stuff, as well as adding general conditions.
	// 	  I'm not sure how useful the conditions will be because if they're already bad they will
	// 	  show up in the alerts

	fmt.Println("Starting TrafficAIO")
	scrape.RunAndSend()

	// merge.GetRelivantTraffic()
}

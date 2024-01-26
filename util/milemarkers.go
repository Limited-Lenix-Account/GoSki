package util

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/twpayne/go-geom/encoding/wkt"
)

func ReadMileMarker() []MileMarker {

	// Opening the File
	file, err := os.Open("data/MILE_MARKERS_GPS.csv")
	if err != nil {
		fmt.Printf("Error Reading Mile File: %s", err)
	}
	defer file.Close()

	//Read the CSV
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error Reading Mile File: %s", err)
	}

	// Loop though the lines and create MileMarker Objects
	var mileMarkers []MileMarker

	for i, record := range records {

		//Ignore Header
		if i == 0 {
			continue
		}

		coord := ParseCoords(record[0])

		t := MileMarker{
			RoadType:    record[5],
			Marker:      record[7],
			Coordinates: coord,
		}
		mileMarkers = append(mileMarkers, t)
	}
	return mileMarkers
}

// Takes a POINT string and returns a Coordinate object (lat, long)
func ParseCoords(raw string) Coordinates {
	geom, err := wkt.Unmarshal(raw)
	if err != nil {
		fmt.Printf("Error Unmarshaling Coordinates!: %s", err)
	}

	coord := geom.FlatCoords()

	t := Coordinates{
		Lat:  coord[1],
		Long: coord[0],
	}
	return t
}

// func ParseRoute(raw string)

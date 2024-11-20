package incidents

import (
	"fmt"

	"github.com/dhconnelly/rtreego"
	"traffic.go/api"
	"traffic.go/internal/plow"
)

func ParseIndidents(tree *rtreego.Rtree) (*[]UsableIncident, error) {
	var incidents []UsableIncident

	fmt.Println("parsing incidnents")
	incResp, err := api.GetIncidents()
	if err != nil {
		return nil, fmt.Errorf("error making api call: %w", err)
	}

	// convert Geometry from interface{} to []float64 or [][]float64
	for _, v := range incResp.Features {
		var usable UsableIncident
		var singlePoint []float64
		var multiPoint [][]float64
		switch geoTypes := v.Geometry.Coordinates.(type) {
		case []interface{}:
			switch geoTypes[0].(type) {
			case float64:
				for _, c := range geoTypes {
					singlePoint = append(singlePoint, c.(float64))
				}

				// convert coordinates to mile marker object
				usable.SinglePoint = rtreego.Point{singlePoint[1], singlePoint[0]}
				usable.SingleMile = plow.FindClosestMileFromPoint(usable.SinglePoint, tree)
			case []interface{}:
				for _, c := range geoTypes {
					var point []float64
					for _, p := range c.([]interface{}) {
						point = append(point, p.(float64))
					}
					multiPoint = append(multiPoint, point)
				}

				// convert coordinates to mile marker object for two points
				usable.PrimaryPoint = rtreego.Point{multiPoint[0][1], multiPoint[0][0]}
				usable.SecondaryPoint = rtreego.Point{multiPoint[1][1], multiPoint[1][0]}
				usable.PrimaryMile = plow.FindClosestMileFromPoint(usable.PrimaryPoint, tree)
				usable.SecondaryMile = plow.FindClosestMileFromPoint(usable.SecondaryPoint, tree)
			default:
				fmt.Println("Something else!", geoTypes)
			}

			incidents = append(incidents, usable)
		}
	}

	return &incidents, nil
}

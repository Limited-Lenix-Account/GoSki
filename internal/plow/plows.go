package plow

import (
	"fmt"

	"traffic.go/api"
)

func ParsePlows() *[]UsePlow {

	var UseablePlows []UsePlow

	plows, err := api.GetSnowPlow()
	if err != nil {
		fmt.Printf("Error making snowplow req %s", err)
	}

	for _, v := range plows.Features {
		var plow UsePlow

		plow.ID = v.AvlLocation.Vehicle.ID
		plow.State = v.AvlLocation.CurrentStatus.State

		p := Point{
			Latitude:  v.AvlLocation.Position.Latitude,
			Longitude: v.AvlLocation.Position.Longitude,
		}

		plow.Position = p

		fmt.Println(v.AvlLocation.CurrentStatus.Info)

		if v.AvlLocation.CurrentStatus.State == "Active" {
			plow.Active = true
		} else {
			plow.Active = false
		}

		UseablePlows = append(UseablePlows, plow)
	}
	return &UseablePlows
}

package plow

import (
	"fmt"
	"sync"

	"github.com/schollz/progressbar/v3"
	"traffic.go/api"
)

func ParsePlows() *[]UsePlow {

	var UseablePlows []UsePlow

	var wg sync.WaitGroup

	apiPlows, _ := api.GetSnowPlowFromAPI("")
	appPlows, err := api.GetSnowPlowFromApp()
	if err != nil {
		fmt.Printf("Error getting APP plows: %s\n", err)
	}
	listLength := len(*appPlows)

	wg.Add(listLength)
	APIBar := progressbar.Default(int64(listLength))
	for i := 0; i < listLength; i++ {
		go func(index int) {
			defer wg.Done()

			res, err := api.GetSnowPlowFromAPI((*appPlows)[index].ID)
			if err != nil {
				fmt.Printf("Error making API req from APP %s\n", err)
			}

			APIBar.Add(1)
			apiPlows.Features = append(apiPlows.Features, res.Features...)

		}(i)
	}

	wg.Wait()

	// for _, v := range *appPlows {
	// 	resp, err := api.GetSnowPlowFromAPI(v.ID)
	// 	APIBar.Add(1)
	// 	if err != nil {
	// 		fmt.Printf("Error making API req from APP %s\n", err)
	// 	}
	// 	apiPlows.Features = append(apiPlows.Features, resp.Features...)
	// }

	if err != nil {
		fmt.Printf("Error making snowplow req %s", err)
	}
	for _, v := range apiPlows.Features {
		var plow UsePlow

		plow.ID = v.AvlLocation.Vehicle.ID
		plow.ID2 = v.AvlLocation.Vehicle.ID2
		plow.State = v.AvlLocation.CurrentStatus.Info

		p := Point{
			Latitude:  v.AvlLocation.Position.Latitude,
			Longitude: v.AvlLocation.Position.Longitude,
		}

		plow.Position = p

		if v.AvlLocation.CurrentStatus.State == "Active" {
			plow.Active = true
		} else {
			plow.Active = false
		}

		UseablePlows = append(UseablePlows, plow)
	}

	return &UseablePlows
}

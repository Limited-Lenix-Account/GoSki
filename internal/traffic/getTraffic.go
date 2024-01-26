package traffic

import (
	"fmt"
	"regexp"
	"strings"

	"traffic.go/api"
)

func ParseTraffic() *[]UseableTraffic {
	var ValidTraffic []UseableTraffic

	travelResp, err := api.GetTravelTimes()
	if err != nil {
		fmt.Println("Error GetTravelTimes()", err)
	}

	for _, v := range travelResp.Features {
		var travelSeg UseableTraffic

		travelSeg.Name = ParseTitle(v.Properties.Name)
		travelSeg.TravelTime = v.Properties.TravelTime
		travelSeg.ID = ParseID(v.Properties.Name)
		travelSeg.COID = v.Properties.ID

		if !(len(v.Properties.SegmentParts) == 0) {
			travelSeg.StartMile = v.Properties.SegmentParts[0].StartMarker
			travelSeg.EndMile = v.Properties.SegmentParts[0].EndMarker
			travelSeg.Route = v.Properties.SegmentParts[0].Route
		}
		travelSeg.UpdatedTime = v.Properties.LastUpdated
		ValidTraffic = append(ValidTraffic, travelSeg)
	}

	return &ValidTraffic
}

func ParseID(raw string) string {

	regexPattern := `(\w+)`
	re := regexp.MustCompile(regexPattern)
	matches := re.FindStringSubmatch(raw)

	return matches[1]

}

func ParseTitle(raw string) string {
	str := strings.Split(raw, " ")[1:]
	res := strings.Join(str, " ")
	return res
}

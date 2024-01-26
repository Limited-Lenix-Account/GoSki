package alerts

import (
	"fmt"
	"regexp"
	"strconv"

	"traffic.go/api"
)

func ParseAlerts() *[]UseableAlert {

	var alertList []UseableAlert

	alertRes, err := api.GetAlerts()
	if err != nil {
		fmt.Printf("Error getting Alert Res %s", err)
	}
	for _, alert := range *alertRes {

		parseAlert := UseableAlert{}

		if alert.AgencyAttribution.AgencyName == "Waze" {
			continue
		} else {
			miles := GetMileIndicators(alert.EventDescription.DescriptionFull)
			switch len(miles) {
			case 0:
				continue
			case 1:
				parseAlert.SingleMile = miles[0]
			case 2:
				parseAlert.StartMile = miles[0]
				parseAlert.EndMile = miles[1]
			default:
				fmt.Println("Unknown Length of Mile List!")
			}

			parseAlert.ID = alert.ID
			parseAlert.Route = alert.Location.RouteDesignator
			parseAlert.BeginTime = alert.BeginTime.Time / 1000
			parseAlert.EndTime = alert.EndTime.Time / 1000
			parseAlert.Tooltip = alert.EventDescription.Tooltip
			parseAlert.Reason = alert.EventDescription.HeadlinePhrase
			parseAlert.Description = alert.EventDescription.CriticalDisruptionHeader

		}

		alertList = append(alertList, parseAlert)
	}
	return &alertList

}

func GetMileIndicators(desc string) []float64 {

	var floatList []float64
	pattern := `Mile Point (\d+(\.\d+)?) to Mile Point (\d+(\.\d+)?)`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(desc)

	for i, match := range matches {

		if (match == "") || (i == 0) {
			continue
		}
		floatVersion, err := strconv.ParseFloat(match, 64)
		if err != nil {
			fmt.Println("error converting str to float", err)
		}
		if floatVersion > 1.0 {
			floatList = append(floatList, floatVersion)
		}
	}

	return floatList
}

func ParseRoute(route string) int {

	return 0
}

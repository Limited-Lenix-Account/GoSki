package merge

import (
	"fmt"
	"strconv"
	"time"

	"traffic.go/internal/alerts"
	"traffic.go/internal/traffic"
)

const (
	LOVELAND_PASS_BEGIN = 216
	LOVELAND_PASS_END   = 230

	VAIL_PASS_BEGIN = 170
	VAIL_PASS_END   = 197

	//Technically all of empire but whatever bruh
	BERTHOUD_PASS_BEGIN = 230
	BERTHOUD_PASS_END   = 260
)

var VALID_COID = []string{
	"OpenTMS-TravelTime7685712394",
	// "OpenTMS-TravelTime554935",
	"OpenTMS-TravelTime7685734533",
	"OpenTMS-TravelTime548989",
}

func GetValidAlerts() (*PassStatus, *PassStatus, *PassStatus) {

	var LovelandPass PassStatus
	var VailPass PassStatus
	var BerthoudPass PassStatus

	alr := alerts.ParseAlerts()
	LovelandPass.Name = "Loveland Pass"
	VailPass.Name = "Vail Pass"
	BerthoudPass.Name = "Berthoud Pass"

	for _, v := range *alr {
		if v.Route == "US 6" {
			if v.StartMile > LOVELAND_PASS_BEGIN && v.EndMile < LOVELAND_PASS_END {
				LovelandPass.Alerts = append(LovelandPass.Alerts, v)
			}

			if !(v.Reason == "Road Closed") {
				LovelandPass.Open = 1
			} else {
				LovelandPass.Open = 0
			}

		} else if v.Route == "I-70" {
			if v.StartMile > VAIL_PASS_BEGIN && v.EndMile < VAIL_PASS_END {
				VailPass.Alerts = append(VailPass.Alerts, v)
			}
			if !(v.Reason == "Road Closed") {
				VailPass.Open = 1
			} else {
				VailPass.Open = 0
			}

		} else if v.Route == "US 40" {
			if v.StartMile > BERTHOUD_PASS_BEGIN && v.EndMile < BERTHOUD_PASS_END {

				fmt.Println(v.Reason)
				BerthoudPass.Alerts = append(BerthoudPass.Alerts, v)
			}

			if !(v.Reason == "Road Closed") {
				BerthoudPass.Open = 1
			} else {
				BerthoudPass.Open = 0
			}

		}
	}
	return &LovelandPass, &VailPass, &BerthoudPass
}

func GetRelivantTraffic() *[]traffic.UseableTraffic {

	var ValidTraffic []traffic.UseableTraffic

	traff := traffic.ParseTraffic()

	for _, v := range *traff {

		if itemInSlice(v.COID, VALID_COID) {
			ValidTraffic = append(ValidTraffic, v)
		}

	}

	PrintTraffic(&ValidTraffic)
	return &ValidTraffic
}

func PrintAlert(li *PassStatus) {

	fmt.Printf("======== %s =======\n", li.Name)

	if len(li.Alerts) == 0 {
		fmt.Println("No Alerts Found ✅")
	}

	for _, t := range li.Alerts {

		mile1 := strconv.FormatFloat(t.StartMile, 'f', -1, 64)
		mile2 := strconv.FormatFloat(t.EndMile, 'f', -1, 64)

		timeStart := time.Unix(t.BeginTime, 0).Format("January 2, 2006, 03:04 PM")
		timeEnd := time.Unix(t.EndTime, 0).Format("January 2, 2006, 03:04 PM")

		fmt.Println(t.Reason)
		fmt.Printf("Affecting From Mile Marker %s to %s\n", mile1, mile2)
		fmt.Printf("From %s to %s\n", timeStart, timeEnd)
		fmt.Println()
	}

}

func PrintTraffic(traff *[]traffic.UseableTraffic) {
	fmt.Println("Current Travel Times:")
	for _, v := range *traff {
		location, _ := time.LoadLocation("America/Denver")
		now := time.Now().In(location)

		diff := now.Sub(v.UpdatedTime)

		fmt.Println(v.Name)
		fmt.Println("Travel Time: ", v.TravelTime/60, "Minutes")
		// fmt.Println(v.UpdatedTime.In(location).Format("January 2, 2006, 03:04 PM"))
		fmt.Println(v.COID)
		fmt.Printf("Last updated %s ago\n", diff)
		fmt.Println()
	}
}

func itemInSlice(item string, slice []string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

package merge

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dhconnelly/rtreego"
	"traffic.go/internal/alerts"
	"traffic.go/internal/incidents"
	"traffic.go/internal/plow"
	"traffic.go/internal/traffic"
)

const (
	LOVELAND_PASS_BEGIN = 216
	LOVELAND_PASS_END   = 230

	VAIL_PASS_BEGIN = 170
	VAIL_PASS_END   = 197

	BERTHOUD_PASS_BEGIN = 230
	BERTHOUD_PASS_END   = 260

	I70_BEING = 100
	I70_END   = 300
)

var VALID_COID = []string{
	"OpenTMS-TravelTime7685712394",
	"OpenTMS-TravelTime7685734533",
	"OpenTMS-TravelTime548989",
}

func Merge(tree *rtreego.Rtree) GrandObject {

	var Total GrandObject

	var Loveland PassStatus
	var Vail PassStatus
	var Berthoud PassStatus

	Loveland.Name = "Loveland Pass"
	Vail.Name = "Vail Pass"
	Berthoud.Name = "Berthoud Pass"

	alr := alerts.ParseAlerts(tree)
	fmt.Println("Getting Alerts...")
	incident, err := incidents.ParseIndidents(tree)
	if err != nil {
		fmt.Println("error getting incidents: %w", err)
	}

	//Get Valid Alerts
	LovelandAlerts, VailAlerts, BerthodAlerts := GetValidAlerts(alr)
	Loveland.Alerts, Vail.Alerts, Berthoud.Alerts = *LovelandAlerts, *VailAlerts, *BerthodAlerts

	IncidentAlerts := GetValidIncidents(incident)

	//Check for Road Closures (These two can be consolidated because they make use of the same request)
	LovelandClosure, VailClosure, BerthoudClosure := GetClosures(alr)
	Loveland.Open, Vail.Open, Berthoud.Open = LovelandClosure, VailClosure, BerthoudClosure

	//Getting Plow Information
	fmt.Println("Getting Plow Information...")
	LovelandPlow, VailPlow, BerthoudPlow := GetSnowPlows(tree)
	Loveland.Plows, Vail.Plows, Berthoud.Plows = *LovelandPlow, *VailPlow, *BerthoudPlow

	//Get Traffic
	fmt.Println("Getting Traffic...")
	traffic := GetRelivantTraffic()
	// Get Incidents
	fmt.Println("Getting incidents...")

	//Build Objects
	Total.LovelandPass = &Loveland
	Total.VailPass = &Vail
	Total.BerthodPass = &Berthoud

	// Incidents will cover like all of I70 and othe relivant roads but doesn't need to be in the same section
	// I don't think
	Total.Traffic = traffic
	Total.Incidents = IncidentAlerts

	return Total

}

func GetValidAlerts(alr *[]alerts.UseableAlert) (*[]alerts.UseableAlert, *[]alerts.UseableAlert, *[]alerts.UseableAlert) {

	var LovelandAlerts []alerts.UseableAlert
	var VailAlerts []alerts.UseableAlert
	var BerthoudAlerts []alerts.UseableAlert

	for _, v := range *alr {
		if v.Route == "US 6" {
			if v.StartMile > LOVELAND_PASS_BEGIN && v.EndMile < LOVELAND_PASS_END {
				LovelandAlerts = append(LovelandAlerts, v)
			}

		} else if v.Route == "I-70" {
			if v.StartMile > VAIL_PASS_BEGIN && v.EndMile < VAIL_PASS_END {
				VailAlerts = append(VailAlerts, v)
			}

		} else if v.Route == "US 40" {
			if v.StartMile > BERTHOUD_PASS_BEGIN && v.EndMile < BERTHOUD_PASS_END {
				BerthoudAlerts = append(BerthoudAlerts, v)
			}
		}
	}
	return &LovelandAlerts, &VailAlerts, &BerthoudAlerts
}

func GetValidIncidents(inc *[]incidents.UsableIncident) *[]incidents.UsableIncident {
	return inc
}

func GetClosures(alr *[]alerts.UseableAlert) (bool, bool, bool) {

	var Loveland, Vail, Berthoud bool
	Berthoud = true
	Vail = true
	Loveland = true

	for _, v := range *alr {
		// fmt.Println(v)
		if v.Route == "US 6" {
			if v.StartMile > LOVELAND_PASS_BEGIN && v.EndMile < LOVELAND_PASS_END {
				if v.Reason == "Road Closed" {
					Loveland = false
				}
			}

		} else if v.Route == "I-70" {
			if v.StartMile > VAIL_PASS_BEGIN && v.EndMile < VAIL_PASS_END {
				if v.Reason == "Road Closed" {
					Vail = false
				}
			}
		} else if v.Route == "US 40" {
			if v.StartMile > BERTHOUD_PASS_BEGIN && v.EndMile < BERTHOUD_PASS_END {
				if v.Reason == "Road Closed" {
					Berthoud = false
				}
			}
		}

	}

	// fmt.Printf("Loveland %t, Vail %t, Berthoud %t\n", Loveland, Vail, Berthoud)
	return Loveland, Vail, Berthoud

}

func GetSnowPlows(tree *rtreego.Rtree) (*[]plow.UsePlow, *[]plow.UsePlow, *[]plow.UsePlow) {
	var LovelandPlow, VailPlow, BerthoudPlow []plow.UsePlow
	plows := plow.DeterminePlowPos(tree)

	for _, v := range *plows {
		if v.ClosestMile != nil {

			if v.ClosestMile.Route == "006F" {
				if v.ClosestMile.Marker > LOVELAND_PASS_BEGIN && v.ClosestMile.Marker < LOVELAND_PASS_END {
					LovelandPlow = append(LovelandPlow, v)
				}

			} else if v.ClosestMile.Route == "070A" {
				// fmt.Println(v.ID, v.ClosestMile.Marker)
				if v.ClosestMile.Marker > VAIL_PASS_BEGIN && v.ClosestMile.Marker < VAIL_PASS_END {
					VailPlow = append(VailPlow, v)
				}

			} else if v.ClosestMile.Route == "040A" {
				if v.ClosestMile.Marker > BERTHOUD_PASS_BEGIN && v.ClosestMile.Marker < BERTHOUD_PASS_END {
					BerthoudPlow = append(BerthoudPlow, v)
				}
			}
		}
	}
	return &LovelandPlow, &VailPlow, &BerthoudPlow
}

func GetRelivantTraffic() *[]traffic.UseableTraffic {

	var ValidTraffic []traffic.UseableTraffic

	traff := traffic.ParseTraffic()

	for _, v := range *traff {

		if itemInSlice(v.COID, VALID_COID) {
			ValidTraffic = append(ValidTraffic, v)
		}

	}
	// PrintTraffic(&ValidTraffic)
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

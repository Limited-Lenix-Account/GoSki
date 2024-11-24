package tg

import (
	"fmt"
	"strings"
	"time"

	"traffic.go/internal/alerts"
	"traffic.go/internal/incidents"
	"traffic.go/internal/merge"
	"traffic.go/internal/plow"
	"traffic.go/internal/traffic"
)

func FormatMessage(Total merge.GrandObject) string {

	lList := AlertToStr(Total.LovelandPass.Alerts)
	lStr := strings.Join(lList, "\n")
	lPlow := SnowPlowStr(Total.LovelandPass.Plows)

	vList := AlertToStr(Total.VailPass.Alerts)
	vStr := strings.Join(vList, "\n")
	vPlow := SnowPlowStr(Total.VailPass.Plows)

	bList := AlertToStr(Total.BerthodPass.Alerts)
	bStr := strings.Join(bList, "\n")
	bPlow := SnowPlowStr(Total.BerthodPass.Plows)

	travelList := TrafficToString(*Total.Traffic)
	travelString := strings.Join(travelList, "\n")

	incidentList := IncidentToStr(*Total.Incidents)
	incidentString := strings.Join(incidentList, "\n")

	finalMessage := []string{
		PassOpen(Total.LovelandPass.Name, Total.LovelandPass.Open),
		lPlow,
		lStr,
		PassOpen(Total.VailPass.Name, Total.VailPass.Open),
		vPlow,
		vStr,
		PassOpen(Total.BerthodPass.Name, Total.BerthodPass.Open),
		bPlow,
		bStr,
		"\n*__Highway Incidents__*\n",
		incidentString,
		"\n*__Some Common Travel Times__*\n",
		travelString,
	}

	testString := strings.Join(finalMessage, "\n")

	return testString

}

func IncidentToStr(inc []incidents.UsableIncident) []string {
	var incList []string

	if len(inc) > 0 {
		incList = append(incList, "*Incidents Found\\!* ⚠️")
		for _, v := range inc {
			incLine := fmt.Sprintf("\n*Reason: _%s_ * \nRoute: %s", RouteToString(v.IncidentType), RouteToString(v.Route))
			incList = append(incList, incLine)

			if v.LanesClosed.LanesStr != "" {
				incStr := RouteToString(v.LanesClosed.LanesStr)
				incList = append(incList, incStr)
			}

		}

	} else {
		incList = append(incList, "*No Incidents Found* *\n")
	}

	return incList
}

func AlertToStr(alr []alerts.UseableAlert) []string {
	var alrList []string

	if len(alr) > 0 {
		alrList = append(alrList, "*Alerts Found\\!* ⚠️\n")
		for _, v := range alr {
			alrLine := fmt.Sprintf("Route: _%s_ \nReason: %s\n", RouteToString(v.Route), v.Reason)
			alrList = append(alrList, alrLine)
		}
	} else {
		alrList = append(alrList, "*No Alerts Found* ☀️\n")
	}
	return alrList
}

func PassOpen(name string, open bool) string {
	var nameStr string
	switch open {
	case false:
		str := fmt.Sprintf("⛔️ *__%s Closed__* ⛔️", name)
		nameStr = str
	case true:
		str := fmt.Sprintf("✅ *__%s Open__* ✅", name)
		nameStr = str
	}
	return nameStr
}

func TrafficToString(traff []traffic.UseableTraffic) []string {

	var trafficList []string

	location, _ := time.LoadLocation("America/Denver")

	for _, v := range traff {
		parsedName := strings.Replace(v.Name, "-", "\\-", -1)
		parsedName = strings.Replace(parsedName, "(", "\\(", -1)
		parsedName = strings.Replace(parsedName, ")", "\\)", -1)
		str := fmt.Sprintf("*%s*\nTravel Time: %d Minutes\nLast Updated: _%s_\n", parsedName, v.TravelTime/60, v.UpdatedTime.In(location).Format("January 2, 2006, 03:04 PM"))
		trafficList = append(trafficList, str)
	}
	// unix := strconv.Itoa(int(time.Now().Unix()))
	timeStamp := fmt.Sprintf("\n\n_Last Update Req: %s_", time.Now().In(location).Format("January 2, 2006, 03:04:05 PM"))
	trafficList = append(trafficList, timeStamp)

	return trafficList
}

func RouteToString(route string) string {
	parsedRoute := strings.Replace(route, "-", "\\-", -1)
	parsedRoute = strings.Replace(parsedRoute, "+", "\\+", -1)
	parsedRoute = strings.Replace(parsedRoute, "|", "\\|", -1)
	return parsedRoute
}

func SnowPlowStr(snow []plow.UsePlow) string {
	snowStr := fmt.Sprintf("_Currently %d Snow Plow\\(s\\) Out_\n", len(snow))
	return snowStr
}

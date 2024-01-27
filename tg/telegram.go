package tg

import (
	"fmt"
	"strings"
	"time"

	"traffic.go/internal/alerts"
	"traffic.go/internal/merge"
	"traffic.go/internal/traffic"
)

func FormatMessage(Total merge.GrandObject) string {

	lList := AlertToStr(Total.LovelandPass.Alerts)
	lStr := strings.Join(lList, "\n")

	vList := AlertToStr(Total.VailPass.Alerts)
	vStr := strings.Join(vList, "\n")

	bList := AlertToStr(Total.BerthodPass.Alerts)
	bStr := strings.Join(bList, "\n")

	travelList := TrafficToString(*Total.Traffic)
	travelString := strings.Join(travelList, "\n")

	finalMessage := []string{
		PassOpen(Total.LovelandPass.Name, Total.LovelandPass.Open),
		lStr,
		PassOpen(Total.VailPass.Name, Total.VailPass.Open),
		vStr,
		PassOpen(Total.BerthodPass.Name, Total.BerthodPass.Open),
		bStr,
		"\n*__Some Common Travel Times__*\n",
		travelString,
	}

	testString := strings.Join(finalMessage, "\n")

	return testString

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

func PassOpen(name string, open int) string {
	var nameStr string
	switch open {
	case 0:
		str := fmt.Sprintf("⛔️ *__%s Closed__* ⛔️\n", name)
		nameStr = str
	case 1:
		str := fmt.Sprintf("✅ *__%s Open__* ✅\n", name)
		nameStr = str
	}
	return nameStr
}

// Work on all this shit again lmaooooo
// Format Traffic List to output to string in tg
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
	return parsedRoute
}

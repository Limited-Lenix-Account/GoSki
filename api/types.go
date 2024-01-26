package api

import "time"

type AlertResponse []struct {
	ID                 string `json:"id"`
	SituationUpdateKey struct {
		SituationID  string `json:"situationId"`
		UpdateNumber int    `json:"updateNumber"`
	} `json:"situationUpdateKey"`
	EventDescription struct {
		DescriptionHeader        string `json:"descriptionHeader"`
		DescriptionBrief         string `json:"descriptionBrief"`
		DescriptionFull          string `json:"descriptionFull"`
		Tooltip                  string `json:"tooltip"`
		CriticalDisruptionHeader string `json:"criticalDisruptionHeader"`
		TellMeText               string `json:"tellMeText"`
		HeadlinePhrase           string `json:"headlinePhrase"`
		LocationDescription      string `json:"locationDescription"`
	} `json:"eventDescription"`
	AgencyAttribution struct {
		AgencyName    string `json:"agencyName"`
		AgencyURL     string `json:"agencyURL"`
		AgencyIconURL string `json:"agencyIconURL"`
	} `json:"agencyAttribution"`
	EditorIdentifier string `json:"editorIdentifier"`
	Icon             struct {
		Image  string `json:"image"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"icon"`
	HeadlinePhrase struct {
		Category int `json:"category"`
		Code     int `json:"code"`
	} `json:"headlinePhrase"`
	Priority int `json:"priority"`
	Bounds   struct {
		MinLat float64 `json:"minLat"`
		MinLon float64 `json:"minLon"`
		MaxLat float64 `json:"maxLat"`
		MaxLon float64 `json:"maxLon"`
	} `json:"bounds"`
	Location struct {
		FipsCode        int    `json:"fipsCode"`
		LinkOwnership   string `json:"linkOwnership"`
		RouteDesignator string `json:"routeDesignator"`
		PrimaryPoint    struct {
			Lat             float64 `json:"lat"`
			Lon             float64 `json:"lon"`
			LinearReference float64 `json:"linearReference"`
		} `json:"primaryPoint"`
		SecondaryPoint struct {
			Lat             float64 `json:"lat"`
			Lon             float64 `json:"lon"`
			LinearReference float64 `json:"linearReference"`
		} `json:"secondaryPoint"`
		LinkDirection        string `json:"linkDirection"`
		RoutePositiveBearing string `json:"routePositiveBearing"`
	} `json:"location"`
	UpdateTime struct {
		Time       int64  `json:"time"`
		TimeZoneID string `json:"timeZoneId"`
		UtcOffset  int    `json:"utcOffset"`
	} `json:"updateTime"`
	BeginTime struct {
		Time       int64  `json:"time"`
		TimeZoneID string `json:"timeZoneId"`
		UtcOffset  int    `json:"utcOffset"`
	} `json:"beginTime"`
	EndTime struct {
		Time       int64  `json:"time"`
		TimeZoneID string `json:"timeZoneId"`
		UtcOffset  int    `json:"utcOffset"`
	} `json:"endTime"`
	TimeZoneID string `json:"timeZoneId"`
	Quantities []any  `json:"quantities"`
	Contacts   []struct {
		OrganizationName string `json:"organizationName"`
	} `json:"contacts"`
	ScheduleOccurrences []struct {
		StartTime struct {
			Time       int64  `json:"time"`
			TimeZoneID string `json:"timeZoneId"`
			UtcOffset  int    `json:"utcOffset"`
		} `json:"startTime"`
		EndTime struct {
			Time       int64  `json:"time"`
			TimeZoneID string `json:"timeZoneId"`
			UtcOffset  int    `json:"utcOffset"`
		} `json:"endTime"`
	} `json:"scheduleOccurrences"`
	LaneImpacts struct {
	} `json:"laneImpacts"`
	Active   bool `json:"active"`
	Verified bool `json:"verified"`
}

type TravelTime struct {
	Type     string `json:"type"`
	Features []struct {
		Type     string `json:"type"`
		Geometry struct {
			Srid        int           `json:"srid"`
			Type        string        `json:"type"`
			Coordinates [][][]float64 `json:"coordinates"`
		} `json:"geometry"`
		Properties struct {
			Name         string    `json:"name"`
			LastUpdated  time.Time `json:"lastUpdated"`
			TravelTime   int       `json:"travelTime"`
			ID           string    `json:"id"`
			SegmentParts []struct {
				Route       string  `json:"route"`
				StartMarker float64 `json:"startMarker"`
				EndMarker   float64 `json:"endMarker"`
			} `json:"segmentParts"`
		} `json:"properties"`
		Attributes struct {
		} `json:"attributes"`
	} `json:"features"`
}

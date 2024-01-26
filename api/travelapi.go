package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	API_KEY = "0JW047K-MNCMYS3-G6R3ZDS-4BHGC0P"
)

func GetTravelTimes() (*TravelTime, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://data.cotrip.org/api/v1/destinations?apiKey="+API_KEY, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "data.cotrip.org")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("dnt", "1")
	req.Header.Set("referer", "https://manage-api.cotrip.org/")
	req.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making travel time request %s\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error Reading travel repsonse%s\n", err)
		return nil, err
	}

	var travelResp TravelTime

	err = json.Unmarshal(body, &travelResp)
	if err != nil {
		fmt.Printf("Error unmarhsalling travel times: %s\n", err)
		return nil, err
	}

	return &travelResp, nil

}

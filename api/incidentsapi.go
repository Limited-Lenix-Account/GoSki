package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetIncidents() (*IncidentsResp, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://data.cotrip.org/api/v1/incidents?apiKey=0JW047K-MNCMYS3-G6R3ZDS-4BHGC0P", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://manage-api.cotrip.org/")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Chromium";v="130", "Google Chrome";v="130", "Not?A_Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var incidentResp IncidentsResp
	err = json.Unmarshal(body, &incidentResp)
	if err != nil {
		fmt.Printf("Error unmarshalling json %s", err)
		return nil, err
	}

	// fmt.Printf("%s\n", body)
	return &incidentResp, nil
}

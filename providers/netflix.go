package providers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type netflix struct{}

// NewNetflix returns a new provider
func NewNetflix() Provider {
	return &netflix{}
}

type netflixAPIResponse struct {
	Positions []struct {
		ID                   int64    `json:"id"`
		Name                 string   `json:"name"`
		Location             string   `json:"location"`
		Locations            []string `json:"locations"`
		Department           string   `json:"department"`
		BusinessUnit         string   `json:"business_unit"`
		AtsJobID             string   `json:"ats_job_id"`
		JobDescription       string   `json:"job_description"`
		WorkLocationOption   string   `json:"work_location_option"`
		CanonicalPositionURL string   `json:"canonicalPositionUrl"`
	} `json:"positions"`
	Count int `json:"count"`
}

func (n *netflix) RetrieveJobs(fn func(job *Job)) error {
	url := "https://explore.jobs.netflix.net/api/apply/v2/jobs?domain=netflix.com&profile=&query=&location="

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var response netflixAPIResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	for _, pos := range response.Positions {
		location := pos.Location
		if len(pos.Locations) > 1 {
			location = strings.Join(pos.Locations, " | ")
		}

		jobType := string(FullTime)
		if pos.WorkLocationOption == "remote" {
			jobType = "Remote"
		}

		fn(&Job{
			Title:    pos.Name,
			Company:  "Netflix",
			Location: location,
			Type:     jobType,
			Desc:     pos.JobDescription,
			Link:     pos.CanonicalPositionURL,
			Misc: map[string]string{
				"department":    pos.Department,
				"business_unit": pos.BusinessUnit,
				"job_id":        pos.AtsJobID,
			},
		})
	}

	return nil
}

package providers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type netflix struct{}

// NewNetflix returns a new provider
func NewNetflix() Provider {
	return &netflix{}
}

type netflixSearch struct {
	RecordCount int `json:"record_count"`
	Records     struct {
		Teams      []interface{} `json:"teams"`
		Categories []interface{} `json:"categories"`
		Locations  []interface{} `json:"locations"`
		Postings   []struct {
			Area               interface{} `json:"area"`
			Text               string      `json:"text"`
			LeverID            string      `json:"lever_id"`
			Team               string      `json:"team"`
			Slug               string      `json:"slug"`
			ExternalID         string      `json:"external_id"`
			Description        string      `json:"description"`
			URL                string      `json:"url"`
			SearchText         string      `json:"search_text"`
			State              string      `json:"state"`
			UpdatedAt          time.Time   `json:"updated_at"`
			CreatedAt          time.Time   `json:"created_at"`
			Location           string      `json:"location"`
			Organization       []string    `json:"organization"`
			LeverTeam          string      `json:"lever_team"`
			AlternateLocations interface{} `json:"alternate_locations"`
			Index              string      `json:"_index"`
			Type               string      `json:"_type"`
			Score              interface{} `json:"_score"`
			Version            interface{} `json:"_version"`
			Explanation        interface{} `json:"_explanation"`
			Sort               interface{} `json:"sort"`
			ID                 string      `json:"id"`
			Highlight          struct {
			} `json:"highlight"`
		} `json:"postings"`
	} `json:"records"`
	Info struct {
		Teams struct {
			Query            string `json:"query"`
			CurrentPage      int    `json:"current_page"`
			NumPages         int    `json:"num_pages"`
			PerPage          int    `json:"per_page"`
			TotalResultCount int    `json:"total_result_count"`
			Facets           struct {
			} `json:"facets"`
		} `json:"teams"`
		Categories struct {
			Query            string `json:"query"`
			CurrentPage      int    `json:"current_page"`
			NumPages         int    `json:"num_pages"`
			PerPage          int    `json:"per_page"`
			TotalResultCount int    `json:"total_result_count"`
			Facets           struct {
			} `json:"facets"`
		} `json:"categories"`
		Locations struct {
			Query            string `json:"query"`
			CurrentPage      int    `json:"current_page"`
			NumPages         int    `json:"num_pages"`
			PerPage          int    `json:"per_page"`
			TotalResultCount int    `json:"total_result_count"`
			Facets           struct {
			} `json:"facets"`
		} `json:"locations"`
		Postings struct {
			Query            string `json:"query"`
			CurrentPage      int    `json:"current_page"`
			NumPages         int    `json:"num_pages"`
			PerPage          int    `json:"per_page"`
			TotalResultCount int    `json:"total_result_count"`
		} `json:"postings"`
	} `json:"info"`
	Errors struct {
	} `json:"errors"`
}

func (netflix *netflix) readPage(url string, fn func(job *Job)) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return handleStatus(res)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	search := netflixSearch{}
	err = json.Unmarshal(body, &search)
	if err != nil {
		return err
	}

	for _, job := range search.Records.Postings {
		fn(&Job{
			Title:    job.Text,
			Company:  "Netflix",
			Location: job.Location,
			Type:     string(FullTime),
			Desc:     job.Description,
			Link:     job.URL,
		})
	}
	return nil
}

func (netflix *netflix) RetrieveJobs(fn func(job *Job)) error {
	res, err := http.Get("https://jobs.netflix.com/api/search")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return handleStatus(res)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	search := netflixSearch{}
	err = json.Unmarshal(body, &search)
	if err != nil {
		return err
	}

	for i := 1; i <= search.Info.Postings.NumPages; i++ {
		err := netflix.readPage("https://jobs.netflix.com/api/search?page="+strconv.Itoa(i), fn)
		if err != nil {
			return err
		}
	}
	return nil
}

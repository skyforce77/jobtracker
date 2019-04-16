package providers

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Netflix struct{}

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
			Facets           struct {
				Location struct {
					LosAngelesCalifornia int `json:"Los Angeles, California"`
					LosGatosCalifornia   int `json:"Los Gatos, California"`
					AmsterdamNetherlands int `json:"Amsterdam, Netherlands"`
					LondonUnitedKingdom  int `json:"London, United Kingdom"`
					SingaporeSingapore   int `json:"Singapore, Singapore"`
					TokyoJapan           int `json:"Tokyo, Japan"`
					AlphavilleBrazil     int `json:"Alphaville, Brazil"`
					SeoulSouthKorea      int `json:"Seoul, South Korea"`
					MumbaiIndia          int `json:"Mumbai, India"`
					MadridSpain          int `json:"Madrid, Spain"`
					ParisFrance          int `json:"Paris, France"`
					SaltLakeCityUtah     int `json:"Salt Lake City, Utah"`
					NewYorkNewYork       int `json:"New York, New York"`
					MakatiPhilippines    int `json:"Makati, Philippines"`
					FremontCalifornia    int `json:"Fremont, California"`
					MexicoCityMexico     int `json:"Mexico City, Mexico"`
				} `json:"location"`
				Team struct {
					Legal                            int `json:"Legal"`
					PostProduction                   int `json:"Post Production"`
					CreativeProduction               int `json:"Creative Production"`
					FinancialPlanningAndAnalysis     int `json:"Financial Planning and Analysis"`
					Marketing                        int `json:"Marketing"`
					Finance                          int `json:"Finance"`
					ContentEngineering               int `json:"Content Engineering"`
					ProductEngineering               int `json:"Product Engineering"`
					Production                       int `json:"Production"`
					Facilities                       int `json:"Facilities"`
					CloudAndPlatformEngineering      int `json:"Cloud and Platform Engineering"`
					PR                               int `json:"PR"`
					DataEngineeringAndInfrastructure int `json:"Data Engineering and Infrastructure"`
					CustomerService                  int `json:"Customer Service"`
					HumanResources                   int `json:"Human Resources"`
					ScienceAndAnalytics              int `json:"Science and Analytics"`
					Content                          int `json:"Content"`
					ProductDesign                    int `json:"Product Design"`
					Recruiting                       int `json:"Recruiting"`
					UIEngineering                    int `json:"UI Engineering"`
					PartnerEcosystemEngineering      int `json:"Partner Ecosystem Engineering"`
					BusinessDevelopment              int `json:"Business Development"`
					ConsumerInsights                 int `json:"Consumer Insights"`
					EmployeeTechnology               int `json:"Employee Technology"`
					InformationSecurity              int `json:"Information Security"`
					StreamingClient                  int `json:"Streaming Client"`
					Globalization                    int `json:"Globalization"`
					ContentDelivery                  int `json:"Content Delivery"`
					CreativeMarketingProduction      int `json:"Creative Marketing Production"`
					PartnerDevices                   int `json:"Partner Devices"`
					ProductManagement                int `json:"Product Management"`
					ContentAcquisition               int `json:"Content Acquisition"`
					CustomerServiceAdministration    int `json:"Customer Service Administration"`
					MarketingProduction              int `json:"Marketing Production"`
					ProductionTechnologies           int `json:"Production Technologies"`
				} `json:"team"`
			} `json:"facets"`
		} `json:"postings"`
	} `json:"info"`
	Errors struct {
	} `json:"errors"`
}

func (netflix *Netflix) readPage(url string) (*list.List, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s\n", res.StatusCode, res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	search := netflixSearch{}
	err = json.Unmarshal(body, &search)
	if err != nil {
		log.Fatal(err)
	}

	jobs := list.New()

	for _, job := range search.Records.Postings {
		jobs.PushBack(&Job{
			Title:    job.Text,
			Company:  "Netflix",
			Location: job.Location,
			Type:     string(FullTime),
			Desc:     job.Description,
			Link:     job.URL,
		})
	}

	return jobs, nil
}

func (netflix *Netflix) ListJobs() *list.List {
	jobs := list.New()

	res, err := http.Get("https://jobs.netflix.com/api/search")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s\n", res.StatusCode, res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	search := netflixSearch{}
	err = json.Unmarshal(body, &search)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= search.Info.Postings.NumPages; i++ {
		j, err := netflix.readPage(fmt.Sprintf("https://jobs.netflix.com/api/search?page=%d", i))
		if err != nil {
			log.Fatal(err)
		}

		jobs.PushBackList(j)
	}

	return jobs
}

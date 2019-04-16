package providers

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Amazon struct{}

const amazonUrl = "https://www.amazon.jobs/fr/search.json?base_query=&city=&country=&county=&facets%5B%5D=location&facets%5B%5D=business_category&facets%5B%5D=category&facets%5B%5D=schedule_type_id&facets%5B%5D=employee_class&facets%5B%5D=normalized_location&facets%5B%5D=job_function_id&latitude=&loc_group_id=&loc_query=&longitude=&query_options=&radius=24km&region=&result_limit=200&sort=recent"

type amazonPage struct {
	Error interface{} `json:"error"`
	Hits  int         `json:"hits"`
	Jobs  []struct {
		BasicQualifications     string `json:"basic_qualifications"`
		BusinessCategory        string `json:"business_category"`
		City                    string `json:"city"`
		CompanyName             string `json:"company_name"`
		CountryCode             string `json:"country_code"`
		Description             string `json:"description"`
		DescriptionShort        string `json:"description_short"`
		ID                      string `json:"id"`
		IDIcims                 string `json:"id_icims"`
		JobCategory             string `json:"job_category"`
		JobPath                 string `json:"job_path"`
		JobScheduleType         string `json:"job_schedule_type"`
		Location                string `json:"location"`
		NormalizedLocation      string `json:"normalized_location"`
		PostedDate              string `json:"posted_date"`
		PreferredQualifications string `json:"preferred_qualifications"`
		PrimarySearchLabel      string `json:"primary_search_label"`
		SourceSystem            string `json:"source_system"`
		State                   string `json:"state"`
		Title                   string `json:"title"`
		UpdatedTime             string `json:"updated_time"`
		URLNextStep             string `json:"url_next_step"`
	} `json:"jobs"`
}

func (amazon *Amazon) ListJobs() *list.List {
	jobs := list.New()

	offset := 0
	hits := 1
	for offset < hits {
		res, err := http.Get(amazonUrl + fmt.Sprintf("&offset=%d", offset))
		if err != nil {
			log.Fatal(err)
		}

		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s\n", res.StatusCode, res.Status)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		search := amazonPage{}
		err = json.Unmarshal(body, &search)
		if err != nil {
			log.Fatal(err)
		}
		if search.Error != nil {
			log.Fatal(search.Error)
		}

		hits = search.Hits
		for _, job := range search.Jobs {
			jobs.PushBack(&Job{
				Title:    job.Title,
				Company:  job.CompanyName,
				Location: job.NormalizedLocation,
				Type:     job.JobScheduleType,
				Desc:     job.Description,
				Link:     "https://www.amazon.jobs" + job.JobPath,
			})
		}

		offset += len(search.Jobs)
		res.Body.Close()
	}

	return jobs
}

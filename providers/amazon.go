package providers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type amazon struct{}

// NewAmazon returns a new provider
func NewAmazon() Provider {
	return &amazon{}
}

const amazonURL = "https://www.amazon.jobs/fr/search.json?base_query=&city=&country=&county=&facets%5B%5D=location&facets%5B%5D=business_category&facets%5B%5D=category&facets%5B%5D=schedule_type_id&facets%5B%5D=employee_class&facets%5B%5D=normalized_location&facets%5B%5D=job_function_id&latitude=&loc_group_id=&loc_query=&longitude=&query_options=&radius=24km&region=&result_limit=200&sort=recent"

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

func (amazon *amazon) RetrieveJobs(fn func(job *Job)) error {
	offset := 0
	hits := 1
	for offset < hits {
		res, err := http.Get(amazonURL + "&offset=" + strconv.Itoa(offset))
		if err != nil {
			return err
		}

		if res.StatusCode != 200 {
			return handleStatus(res)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		search := amazonPage{}
		err = json.Unmarshal(body, &search)
		if err != nil {
			return err
		}
		if search.Error != nil {
			return errors.New("amazon err request")
		}

		hits = search.Hits
		for _, job := range search.Jobs {
			fn(&Job{
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
	return nil
}

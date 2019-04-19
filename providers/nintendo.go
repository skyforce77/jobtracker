package providers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type nintendo struct{}

func NewNintendo() *nintendo {
	return &nintendo{}
}

type nintendoSearch []struct {
	JobID                     string `json:"JobId"`
	JobTitle                  string `json:"JobTitle"`
	JobDescription            string `json:"JobDescription"`
	JobPrimaryLocationCode    string `json:"JobPrimaryLocationCode"`
	JobLocationState          string `json:"JobLocationState"`
	JobLocationStateAbbrev    string `json:"JobLocationStateAbbrev"`
	DescriptionExternalHTML   string `json:"DescriptionExternalHTML"`
	ExternalQualificationHTML string `json:"ExternalQualificationHTML"`
	JobCreationDate           string `json:"JobCreationDate"`
}

func (nintendo *nintendo) RetrieveJobs(fn func(job *Job)) {
	res, err := http.Get("https://2oc84v7py6.execute-api.us-west-2.amazonaws.com/prod/api/jobs/")
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

	search := nintendoSearch{}
	err = json.Unmarshal(body, &search)
	if err != nil {
		log.Fatal(err)
	}

	for _, j := range search {
		fn(&Job{
			j.JobTitle,
			"Nintendo",
			j.JobPrimaryLocationCode+","+j.JobLocationState,
			string(FullTime),
			j.DescriptionExternalHTML,
			"https://careers.nintendo.com/job-openings/listing/"+j.JobID+".html",
			nil,
		})
	}
}

package providers

import (
	"encoding/json"
	"github.com/k3a/html2text"
	"io/ioutil"
	"net/http"
	"strings"
)

type nintendo struct{}

// NewNintendo returns a new provider
func NewNintendo() Provider {
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

func (nintendo *nintendo) RetrieveJobs(fn func(job *Job)) error {
	res, err := http.Get("https://2oc84v7py6.execute-api.us-west-2.amazonaws.com/prod/api/jobs/")
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

	search := nintendoSearch{}
	err = json.Unmarshal(body, &search)
	if err != nil {
		return err
	}

	for _, j := range search {
		fn(&Job{
			j.JobTitle,
			"Nintendo",
			j.JobPrimaryLocationCode + "," + j.JobLocationState,
			string(FullTime),
			strings.TrimSpace(html2text.HTML2Text(j.DescriptionExternalHTML)),
			"https://careers.nintendo.com/job-openings/listing/" + j.JobID + ".html",
			nil,
		})
	}
	return nil
}

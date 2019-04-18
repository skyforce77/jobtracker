package providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type lever struct {
	company string
	slug    string
}

func NewLever() *lever {
	return &lever{
		"Lever",
		"lever",
	}
}

type leverSearch []struct {
	Title    string `json:"title"`
	Postings []struct {
		Additional      string `json:"additional"`
		AdditionalPlain string `json:"additionalPlain"`
		Categories      struct {
			Commitment string `json:"commitment"`
			Department string `json:"department"`
			Location   string `json:"location"`
			Team       string `json:"team"`
		} `json:"categories"`
		CreatedAt        int64  `json:"createdAt"`
		DescriptionPlain string `json:"descriptionPlain"`
		Description      string `json:"description"`
		ID               string `json:"id"`
		Lists            []struct {
			Text    string `json:"text"`
			Content string `json:"content"`
		} `json:"lists"`
		Text      string `json:"text"`
		HostedURL string `json:"hostedUrl"`
		ApplyURL  string `json:"applyUrl"`
	} `json:"postings"`
}

func (lever *lever) RetrieveJobs(fn func(job *Job)) {
	res, err := http.Get(fmt.Sprintf("https://api.lever.co/v0/postings/%s?group=team&mode=json", lever.slug))
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
	search := leverSearch{}
	err = json.Unmarshal(body, &search)
	if err != nil {
		log.Fatal(err)
	}

	for _, category := range search {
		for _, job := range category.Postings {
			fn(&Job{
				Title:    job.Text,
				Company:  lever.company,
				Location: job.Categories.Location,
				Type:     job.Categories.Commitment,
				Desc:     job.Description,
				Link:     job.HostedURL,
			})
		}
	}

	res.Body.Close()
}

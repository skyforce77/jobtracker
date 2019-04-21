package providers

import (
	"encoding/json"
	"github.com/k3a/html2text"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type lever struct {
	company string
	slug    string
}

// NewLever returns a new provider
func NewLever() Provider {
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

func (lever *lever) RetrieveJobs(fn func(job *Job)) error {
	res, err := http.Get("https://api.lever.co/v0/postings/" + lever.slug + "?group=team&mode=json")
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("status code error:" + strconv.Itoa(res.StatusCode) + " " + res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	search := leverSearch{}
	err = json.Unmarshal(body, &search)
	if err != nil {
		return err
	}

	for _, category := range search {
		for _, job := range category.Postings {
			fn(&Job{
				Title:    job.Text,
				Company:  lever.company,
				Location: job.Categories.Location,
				Type:     job.Categories.Commitment,
				Desc:     strings.TrimSpace(html2text.HTML2Text(job.Description)),
				Link:     job.HostedURL,
			})
		}
	}

	res.Body.Close()
	return nil
}

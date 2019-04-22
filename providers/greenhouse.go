package providers

import (
	"encoding/json"
	"errors"
	"github.com/k3a/html2text"
	"html"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type greenhouse struct {
	company string
	slug    string
}

// NewGreenhouse returns a new provider
func NewGreenhouse() Provider {
	return &greenhouse{
		"Greenhouse",
		"greenhouse",
	}
}

type greenhouseSearch struct {
	Jobs []struct {
		AbsoluteURL   string `json:"absolute_url"`
		Location      struct {
			Name string `json:"name"`
		} `json:"location"`
		ID            int         `json:"id"`
		Title         string      `json:"title"`
	} `json:"jobs"`
}

type greenhouseJob struct {
	Content       string      `json:"content"`
}

func (greenhouse *greenhouse) RetrieveJob(job *Job, jobId int, fn func(job *Job)) error {
	res, err := http.Get("https://api.greenhouse.io/v1/boards/" +
		greenhouse.slug + "/embed/job/?id=" + strconv.Itoa(jobId))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New("status code error:" + strconv.Itoa(res.StatusCode) + " " + res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	search := greenhouseJob{}
	err = json.Unmarshal(body, &search)
	if err != nil {
		return err
	}

	job.Desc = html2text.HTML2Text(
		strings.TrimSpace(html.UnescapeString(search.Content)))
	fn(job)
	return nil
}

func (greenhouse *greenhouse) RetrieveJobs(fn func(job *Job)) error {
	res, err := http.Get("https://api.greenhouse.io/v1/boards/" + greenhouse.slug + "/jobs")
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
	search := greenhouseSearch{}
	err = json.Unmarshal(body, &search)
	if err != nil {
		return err
	}

	for _, job := range search.Jobs {
		j := &Job{
			Title:    job.Title,
			Company:  greenhouse.company,
			Location: job.Location.Name,
			Type:     string(FullTime),
			Link:     job.AbsoluteURL,
		}

		err := greenhouse.RetrieveJob(j, job.ID, fn)
		if err != nil {
			return err
		}
	}

	res.Body.Close()
	return nil
}

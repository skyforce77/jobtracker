package providers

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type Whittard struct{}

func (whittard *Whittard) requestJob(url string) (*Job, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	job := Job{
		Link:    url,
		Company: "Whittard",
	}

	doc.Find(".position_title .jobs-row-input").Each(func(i int, s *goquery.Selection) {
		job.Title = s.Text()
	})
	doc.Find(".position_description .jobs-row-input p").Each(func(i int, s *goquery.Selection) {
		job.Desc += s.Text()
	})
	doc.Find(".position_employment_type .jobs-row-input").Each(func(i int, s *goquery.Selection) {
		job.Type = s.Text()
	})
	doc.Find(".position_job_location .jobs-row-input").Each(func(i int, s *goquery.Selection) {
		job.Location = s.Text()
	})

	return &job, nil
}

func (whittard *Whittard) ListJobs() []*Job {
	jobs := make([]*Job, 0)

	res, err := http.Get("https://careers.whittard.co.uk/contact/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s\n", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".job-cta a").Each(func(i int, s *goquery.Selection) {
		url, ok := s.Attr("href")
		if !ok {
			log.Fatal(err)
		}

		job, err := whittard.requestJob(url)
		if err != nil {
			log.Fatal(err)
		}

		jobs = append(jobs, job)
	})

	return jobs
}

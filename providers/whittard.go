package providers

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type whittard struct{}

func NewWhittard() *whittard {
	return &whittard{}
}

func (whittard *whittard) requestJob(url string, fn func(job *Job)) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
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

	fn(&job)
	return nil
}

func (whittard *whittard) RetrieveJobs(fn func(job *Job)) error {
	res, err := http.Get("https://careers.whittard.co.uk/contact/")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return handleStatus(res)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	doc.Find(".job-cta a").Each(func(i int, s *goquery.Selection) {
		url, ok := s.Attr("href")
		if !ok {
			return
		}

		err = whittard.requestJob(url, fn)
	})
	return err
}

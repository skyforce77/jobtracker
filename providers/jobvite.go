package providers

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type jobVite struct {
	company string
	url     string
}

func (jobvite *jobVite) requestJob(job *Job, fn func(job *Job)) error {
	res, err := http.Get(job.Link)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	doc.Find("h2.jv-header").First().Each(func(i int, s *goquery.Selection) {
		job.Title = s.Text()
	})
	doc.Find("div.jv-job-detail-description").Each(func(i int, s *goquery.Selection) {
		job.Desc = s.Text()
	})

	fn(job)
}

func (jobvite *jobVite) RetrieveJobs(fn func(job *Job)) error {
	res, err := http.Get(jobvite.url)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return HandleStatus(res)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	doc.Find(".jv-job-list tbody tr").Each(func(j int, s *goquery.Selection) {
		job := Job{
			Company: jobvite.company,
			Type:    string(FullTime),
		}

		s.Children().Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				job.Title = s.Children().First().Text()
				url, ok := s.Children().First().Attr("href")
				if !ok {
					return
				}

				job.Link = "https://jobs.jobvite.com" + url
			} else if i == 1 {
				job.Location = s.Text()
			}
		})

		jobvite.requestJob(&job, fn)
	})

	doc.Find(".jv-job-list ul li a").Each(func(j int, s *goquery.Selection) {
		url, ok := s.Attr("href")
		if !ok {
			return
		}

		job := Job{
			Company: jobvite.company,
			Type:    string(FullTime),
			Link:    "https://jobs.jobvite.com" + url,
		}

		s.Children().Each(func(i int, s *goquery.Selection) {
			if i == 1 {
				job.Location = s.Text()
			}
		})

		err = jobvite.requestJob(&job, fn)
		if err != nil {
			println(err)
		}
	})

	res.Body.Close()
	return nil
}

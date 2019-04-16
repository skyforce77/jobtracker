package providers

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type twitter struct{}

func NewTwitter() *twitter {
	return &twitter{}
}

const twitterUrl = "https://careers.twitter.com/content/careers-twitter/en/jobs-search.html?q=&team=&location="

func (twitter *twitter) requestJob(url string, fn func(job *Job)) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	job := Job{
		Link:    url,
		Company: "Twitter",
		Type: string(FullTime),
	}

	doc.Find("h1").First().Each(func(i int, s *goquery.Selection) {
		job.Title = s.Text()
	})
	doc.Find("#main-content div div div div div").Each(func(i int, s *goquery.Selection) {
		if i == 1 {
			job.Location = s.Text()
		} else if i == 2 {
			job.Desc = s.Text()
		}
	})

	fn(&job)
}

func (twitter *twitter) RetrieveJobs(fn func(job *Job)) {
	i := 0
	ni := -1
	for ni != i{
		ni = i

		res, err := http.Get(twitterUrl + fmt.Sprintf("&start=%d", i))
		if err != nil {
			log.Fatal(err)
		}

		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s\n", res.StatusCode, res.Status)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find(".job-search-entries .job-search-item a").Each(func(j int, s *goquery.Selection) {
			url, ok := s.Attr("href")
			if !ok {
				log.Fatal(err)
			}

			i++
			twitter.requestJob("https://careers.twitter.com" + url, fn)
		})

		res.Body.Close()
	}
}

package providers

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type disney struct{}

func NewDisney() *disney {
	return &disney{}
}

type disneySearch struct {
	Filters string `json:"results"`
	HasJobs bool   `json:"hasJobs"`
}

const disneyUrl = "https://jobs.disneycareers.com/search-jobs/results?ActiveFacetID=0&CurrentPage=%d&RecordsPerPage=15&Distance=50&RadiusUnitType=0&Keywords=&Location=&Latitude=&Longitude=&ShowRadius=False&CustomFacetName=&FacetTerm=&FacetType=0&SearchResultsModuleName=Search+Results&SearchFiltersModuleName=Search+Filters&SortCriteria=0&SortDirection=1&SearchType=5&CategoryFacetTerm=&CategoryFacetType=&LocationFacetTerm=&LocationFacetType=&KeywordType=&LocationType=&LocationPath=&OrganizationIds=&PostalCode=&fc=&fl=&fcf=&afc=&afl=&afcf="

func (disney *disney) requestJob(url string, fn func(job *Job)) {
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
		Link: url,
		Type: string(FullTime),
	}

	doc.Find("#job-title-scrape").Each(func(i int, s *goquery.Selection) {
		job.Title = s.Text()
	})
	doc.Find(".ats-description").Each(func(i int, s *goquery.Selection) {
		job.Desc += s.Text()
	})
	doc.Find("#job-brand").Children().Remove().End().Each(func(i int, s *goquery.Selection) {
		job.Company = s.Text()
	})
	doc.Find(".job-location-scrape").Each(func(i int, s *goquery.Selection) {
		job.Location = s.Text()
	})

	fn(&job)
}

func (disney *disney) readPage(page int, search *disneySearch, fn func(job *Job)) {
	res, err := http.Get(fmt.Sprintf(disneyUrl, page))
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

	err = json.Unmarshal(body, search)
	if err != nil {
		log.Fatal(err)
	}

	rdr := strings.NewReader(search.Filters)
	doc, err := goquery.NewDocumentFromReader(rdr)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		url, ok := s.Attr("href")
		if !ok {
			log.Fatal(err)
		}

		if strings.HasPrefix(url, "/job/") {
			disney.requestJob("https://jobs.disneycareers.com"+url, fn)
		}
	})
}

func (disney *disney) RetrieveJobs(fn func(job *Job)) {
	search := &disneySearch{
		"",
		true,
	}

	i := 0
	for search.HasJobs {
		i++

		disney.readPage(i, search, fn)
	}
}

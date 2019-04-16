package providers

import (
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Disney struct{}

type disneySearch struct {
	Filters string `json:"results"`
	HasJobs bool   `json:"hasJobs"`
}

const disneyUrl = "https://jobs.disneycareers.com/search-jobs/results?ActiveFacetID=0&CurrentPage=%d&RecordsPerPage=15&Distance=50&RadiusUnitType=0&Keywords=&Location=&Latitude=&Longitude=&ShowRadius=False&CustomFacetName=&FacetTerm=&FacetType=0&SearchResultsModuleName=Search+Results&SearchFiltersModuleName=Search+Filters&SortCriteria=0&SortDirection=1&SearchType=5&CategoryFacetTerm=&CategoryFacetType=&LocationFacetTerm=&LocationFacetType=&KeywordType=&LocationType=&LocationPath=&OrganizationIds=&PostalCode=&fc=&fl=&fcf=&afc=&afl=&afcf="

func (disney *Disney) requestJob(url string) (*Job, error) {
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

	return &job, nil
}

func (disney *Disney) readPage(page int, search *disneySearch) *list.List {
	jobs := list.New()

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
			j, err := disney.requestJob("https://jobs.disneycareers.com" + url)
			if err != nil {
				log.Fatal(err)
			}
			jobs.PushBack(j)
		}
	})

	return jobs
}

func (disney *Disney) ListJobs() *list.List {
	jobs := list.New()

	search := &disneySearch{
		"",
		true,
	}

	i := 0
	for search.HasJobs {
		i++

		j := disney.readPage(i, search)
		jobs.PushBackList(j)
	}

	return jobs
}

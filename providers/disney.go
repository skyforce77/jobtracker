package providers

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type disney struct{}

// NewDisney returns a new provider
func NewDisney() Provider {
	return &disney{}
}

type disneySearch struct {
	Filters string `json:"results"`
	HasJobs bool   `json:"hasJobs"`
}

const disneyURL = "https://jobs.disneycareers.com/search-jobs/results?ActiveFacetID=0&RecordsPerPage=15&Distance=50&RadiusUnitType=0&Keywords=&Location=&Latitude=&Longitude=&ShowRadius=False&CustomFacetName=&FacetTerm=&FacetType=0&SearchResultsModuleName=Search+Results&SearchFiltersModuleName=Search+Filters&SortCriteria=0&SortDirection=1&SearchType=5&CategoryFacetTerm=&CategoryFacetType=&LocationFacetTerm=&LocationFacetType=&KeywordType=&LocationType=&LocationPath=&OrganizationIds=&PostalCode=&fc=&fl=&fcf=&afc=&afl=&afcf=&CurrentPage="

func (disney *disney) requestJob(url string, fn func(job *Job)) error {
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
	return nil
}

func (disney *disney) readPage(page int, search *disneySearch, fn func(job *Job)) error {
	res, err := http.Get(disneyURL + strconv.Itoa(page))
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

	err = json.Unmarshal(body, search)
	if err != nil {
		return err
	}

	rdr := strings.NewReader(search.Filters)
	doc, err := goquery.NewDocumentFromReader(rdr)
	if err != nil {
		return err
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		url, ok := s.Attr("href")
		if !ok {
			return
		}

		if strings.HasPrefix(url, "/job/") {
			err = disney.requestJob("https://jobs.disneycareers.com"+url, fn)
		}
	})
	return err
}

func (disney *disney) RetrieveJobs(fn func(job *Job)) error {
	search := &disneySearch{
		"",
		true,
	}

	i := 0
	for search.HasJobs {
		i++

		err := disney.readPage(i, search, fn)
		if err != nil {
			return err
		}
	}
	return nil
}

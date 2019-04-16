package providers

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type pokemon struct{}

func NewPokemon() *pokemon {
	return &pokemon{}
}

const pokemonUrl = "https://chj.tbe.taleo.net/chj04/ats/careers/searchResults.jsp?org=POKEMON&cws=1"

func (pokemon *pokemon) requestJob(url string, fn func(job *Job)) {
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
		Company: "Pokémon",
		Type:    string(FullTime),
	}

	doc.Find("table tbody tr td h1").First().Each(func(i int, s *goquery.Selection) {
		job.Title = s.Text()
	})
	doc.Find("table tbody tr td b").First().Each(func(i int, s *goquery.Selection) {
		job.Location = s.Text()
	})
	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		if i == 8 {
			job.Desc = s.Text()
		}
	})

	fn(&job)
}

func (pokemon *pokemon) RetrieveJobs(fn func(job *Job)) {
	res, err := http.Get(pokemonUrl)
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

	doc.Find("#cws-search-results tbody tr td b a").Each(func(i int, s *goquery.Selection) {
		url, ok := s.Attr("href")
		if !ok {
			log.Fatal(err)
		}

		pokemon.requestJob(url, fn)
	})
}

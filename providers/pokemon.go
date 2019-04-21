package providers

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type pokemon struct{}

// NewPokemon returns a new provider
func NewPokemon() Provider {
	return &pokemon{}
}

const pokemonURL = "https://chj.tbe.taleo.net/chj04/ats/careers/searchResults.jsp?org=POKEMON&cws=1"

func (pokemon *pokemon) requestJob(url string, fn func(job *Job)) error {
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
	return nil
}

func (pokemon *pokemon) RetrieveJobs(fn func(job *Job)) error {
	res, err := http.Get(pokemonURL)
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

	doc.Find("#cws-search-results tbody tr td b a").Each(func(i int, s *goquery.Selection) {
		url, ok := s.Attr("href")
		if !ok {
			return
		}

		err = pokemon.requestJob(url, fn)
	})
	return err
}

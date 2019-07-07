package providers

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type civiweb struct {
	latest bool
}

// NewCiviweb returns a new provider
func NewCiviweb() Provider {
	return &civiweb{
		false,
	}
}

type civiwebLatest struct {
	civiweb
}

// NewCiviwebLatest returns a new provider
func NewCiviwebLatest() Provider {
	return &civiwebLatest{
		civiweb{true},
	}
}

const civiwebURL = "https://www.civiweb.com/FR/offre-liste/page/1.aspx"

func (civiweb *civiweb) findPageCount() (int, error) {
	res, err := http.Get("https://www.civiweb.com/FR/offre-liste/page/1.aspx")
	if err != nil {
		return 0, err
	}

	if res.StatusCode != 200 {
		return 0, handleStatus(res)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return 0, err
	}

	regex, err := regexp.Compile("/FR/offre-liste/page/([0-9]+).aspx")
	if err != nil {
		return 0, err
	}

	i := 0
	doc.Find("a").Each(func(j int, s *goquery.Selection) {
		url, ok := s.Attr("href")
		if !ok {
			return
		}

		m := regex.FindAllStringSubmatch(url, -1)
		for _, n := range m {
			for _, o := range n {
				a, err := strconv.Atoi(o)
				if err == nil && i < a {
					i = a
				}
			}
		}
	})
	return i, nil
}

func (civiweb *civiweb) retrieveJob(fn func(job *Job), title string,
	url string, regex *regexp.Regexp) error {

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return handleStatus(res)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	job := &Job{
		Title: title,
		Type:  string(VI),
		Link:  url,
		Misc:  make(map[string]string),
	}

	intro := doc.Find(".intro-offre").First().Text()
	text := doc.Find(".txt-offre").First().Text()
	matches := regex.FindAllStringSubmatch(intro, -1)
	if len(matches) > 0 {
		m := matches[0]

		job.Desc = text
		job.Location = m[2] + ", " + m[1]
		job.Company = m[7]

		job.Misc["start"] = m[3]
		job.Misc["end"] = m[4]
		job.Misc["duration"] = m[5] + " " + m[6]
		job.Misc["salary"] = m[8]

		fn(job)
	}
	return nil
}

func (civiweb *civiweb) retrieveLastJobs(fn func(job *Job), regex *regexp.Regexp) error {
	res, err := http.Get("https://www.civiweb.com/FR/index.aspx")
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return handleStatus(res)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	doc.Find(".last-offers ul li a").Each(func(j int, s *goquery.Selection) {
		url, ok := s.Attr("href")
		if !ok {
			return
		}

		err := civiweb.retrieveJob(fn, strings.TrimSpace(strings.Split(s.Text(), "\n")[1]),
			"https://www.civiweb.com"+url, regex)
		if err != nil {
			panic(err)
		}
	})
	return nil
}

func (civiweb *civiweb) RetrieveJobs(fn func(job *Job)) error {
	detailsLinkRegex, err := regexp.Compile(`\s*(.+)\n +\((.+)\)\n +du\n +(.+)\n +au\n +(.+)\n +\(pour\n +(.+)\n +(.+)\)\n +ETABLISSEMENT :\n +(.+)\n +REMUNERATION MENSUELLE :\n +(.+)€`)
	if err != nil {
		return err
	}

	if civiweb.latest {
		return civiweb.retrieveLastJobs(fn, detailsLinkRegex)
	}

	count, err := civiweb.findPageCount()
	if err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		res, err := http.Get("https://www.civiweb.com/FR/offre-liste/page/" + strconv.Itoa(i) + ".aspx")
		if err != nil {
			return err
		}

		if res.StatusCode != 200 {
			return handleStatus(res)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return err
		}

		doc.Find("#content2 a.xt_offrelink").Each(func(j int, s *goquery.Selection) {
			url, ok := s.Attr("href")
			if !ok {
				return
			}

			err := civiweb.retrieveJob(fn, s.Text(),
				"https://www.civiweb.com"+url, detailsLinkRegex)
			if err != nil {
				panic(err)
			}
		})

		res.Body.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

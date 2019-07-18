package print

import (
	"../providers"
	"../util"
	"container/list"
	"strings"
	"unicode/utf8"
	"fmt"
)

// Print will print jobs
func Print(pro []providers.Provider, pretty bool, filter *string) {
	if !pretty {
		for _, p := range pro {
			if filter != nil {
				p.RetrieveJobs(util.Filter(*filter, func(job *providers.Job) {
					fmt.Printf("\"%s\",\"%s\",\"%s\",\"%s\",\"%s\"\n", job.Title, job.Company, job.Location, job.Type, job.Link)
				}))
			} else {
				p.RetrieveJobs(func(job *providers.Job) {
					fmt.Printf("\"%s\",\"%s\",\"%s\",\"%s\",\"%s\"\n", job.Title, job.Company, job.Location, job.Type, job.Link)
				})
			}
			return
		}
	}

	lst := list.New()
	for _, p := range pro {
		if filter != nil {
			lst.PushBackList(util.FilterCollect(p, *filter))
		} else {
			lst.PushBackList(providers.Collect(p))
		}
	}

	if lst.Len() == 0 {
		return
	}

	max := []int{0, 0, 0, 0}
	providers.IterateOver(lst, func(job *providers.Job) {
		if max[0] < utf8.RuneCountInString(job.Title) {
			max[0] = utf8.RuneCountInString(job.Title)
		}
		if max[1] < utf8.RuneCountInString(job.Company) {
			max[1] = utf8.RuneCountInString(job.Company)
		}
		if max[2] < utf8.RuneCountInString(job.Location) {
			max[2] = utf8.RuneCountInString(job.Location)
		}
		if max[3] < utf8.RuneCountInString(job.Link) {
			max[3] = utf8.RuneCountInString(job.Link)
		}
	})

	for i := 0; i < len(max); i++ {
		max[i] += 1
	}

	prettyPrint("Title", max[0])
	prettyPrint("Company", max[1])
	prettyPrint("Location", max[2])
	prettyPrint("Link", max[3])
	println()

	providers.IterateOver(lst, func(job *providers.Job) {
		prettyPrint(job.Title, max[0])
		prettyPrint(job.Company, max[1])
		prettyPrint(job.Location, max[2])
		prettyPrint(job.Link, max[3])
		println()
	})
}

func prettyPrint(txt string, ln int) {
	lnt := utf8.RuneCountInString(txt)
	if ln < lnt {
		ln = lnt+1
	}
	print(txt)
	print(strings.Repeat(" ", ln-lnt))
}

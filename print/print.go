package print

import (
	"../providers"
	"container/list"
	"strings"
	"unicode/utf8"
)

// Print will print jobs
func Print(pro []providers.Provider, pretty bool) {
	if !pretty {
		for _, p := range pro {
			p.RetrieveJobs(func(job *providers.Job) {
				println(job.Title, "|", job.Company, "|", job.Location, "|", job.Link)
			})
			return
		}
	}

	lst := list.New()
	for _, p := range pro {
		lst.PushBackList(providers.Collect(p))
	}

	max := []int{0, 0, 0, 0}
	providers.IterateOver(lst, func(job *providers.Job) {
		if max[0] < len(job.Title) {
			max[0] = len(job.Title)
		}
		if max[1] < len(job.Company) {
			max[1] = len(job.Company)
		}
		if max[2] < len(job.Location) {
			max[2] = len(job.Location)
		}
		if max[3] < len(job.Link) {
			max[3] = len(job.Link)
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
	print(txt)
	print(strings.Repeat(" ", ln-utf8.RuneCountInString(txt)))
}

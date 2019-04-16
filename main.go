package main

import (
	"./providers"
	"log"
)

func main() {
	p := providers.Netflix{}
	lst := p.ListJobs()

	providers.IterateOver(lst, func(job *providers.Job) {
		log.Println(job.Title, job.Company)
	})
}

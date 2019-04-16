package main

import (
	"./providers"
	"log"
)

func main() {
	p := providers.NewPaloAltoNetworks()
	p.RetrieveJobs(func(job *providers.Job) {
		log.Println(job.Title, job.Company)
	})
}

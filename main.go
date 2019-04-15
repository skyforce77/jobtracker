package main

import (
	"./providers"
	"log"
)

func main() {
	p := providers.Disney{}
	for _, v := range p.ListJobs() {
		log.Println(v.Title, v.Company)
	}
}

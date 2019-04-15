package main

import (
	"./providers"
	"log"
)

func main() {
	p := providers.Netflix{}
	for _, v := range p.ListJobs() {
		log.Println(v.Title)
	}
}

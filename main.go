package main

import (
	"./providers"
	"log"
)

func main() {
	w := providers.Whittard{}
	for _, v := range w.ListJobs() {
		log.Println(v)
	}
}
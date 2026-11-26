package main

import (
	"log"

	"github.com/daniilchuits/pre-trainee-practice/internal/report"
)

func main() {

	recs := []report.Record{{SKU: "qwe", Qty: 1, PriceCents: 5}}
	_ = recs

	j1 := report.JSONExporter{}

	b, err := j1.Export(nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(b))
}

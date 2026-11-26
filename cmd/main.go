package main

import (
	"fmt"
	"log"

	"github.com/daniilchuits/pre-trainee-practice/internal/report"
)

func main() {

	recs := []report.Record{{SKU: "A-1", Qty: 2, PriceCents: 199}, {SKU: "B-9", Qty: 1, PriceCents: 50}}
	_ = recs

	j := report.JSONExporter{}
	b, err := j.Export(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	c := report.CSVExporter{}
	b, err = c.Export(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

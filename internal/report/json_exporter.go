package report

import (
	"encoding/json"
	"fmt"
)

type jsonRecord struct {
	SKU        string `json:"sku"`
	Qty        int    `json:"qty"`
	PriceCents int    `json:"price_cents"`
	LineCents  int    `json:"line_cents"`
}

func (JSONExporter) Name() string {
	return "json"
}

func (JSONExporter) Export(records []Record) ([]byte, error) {

	var jsonRecords []jsonRecord

	for i := range records {
		jsonRecords = append(jsonRecords, jsonRecord{
			SKU:        records[i].SKU,
			Qty:        records[i].Qty,
			PriceCents: records[i].PriceCents,
			LineCents:  records[i].PriceCents * records[i].Qty,
		})
	}

	bytes, err := json.Marshal(jsonRecords)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", JsonMarshaling, err)
	}
	return bytes, nil
}

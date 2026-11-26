package report

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
)

func (CSVExporter) Name() string {
	return "csv"
}

func (c CSVExporter) Export(records []Record) ([]byte, error) {

	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	sep := c.separator()

	if err := w.Write([]string{
		"sku",
		"qty",
		"price_cents",
		"line_cents",
	}); err != nil {
		return nil, fmt.Errorf("%w: %w", CsvMarshling, err)
	}

	w.Comma = sep

	for i := range records {
		lineCents, err := recordValidation(records[i])
		if err != nil {
			return nil, fmt.Errorf(
				"%w | record №%d",
				err, i+1,
			)
		}
		w.Write([]string{
			records[i].SKU,
			strconv.Itoa(records[i].Qty),
			strconv.Itoa(records[i].PriceCents),
			lineCents,
		},
		)
	}
	w.Flush()

	if err := w.Error(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c CSVExporter) separator() rune {
	if c.Comma == 0 {
		return ','
	}
	return c.Comma
}

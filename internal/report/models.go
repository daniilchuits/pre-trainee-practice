package report

import "errors"

type Record struct {
	SKU        string
	Qty        int
	PriceCents int // unit price
}

type Exporter interface {
	Name() string
	Export(records []Record) ([]byte, error)
}

type JSONExporter struct{}

type CSVExporter struct {
	Comma rune // if 0, use ','
}

var (
	JsonMarshaling error = errors.New("Error marshaling json")
	CsvMarshling   error = errors.New("Error marshaling csv")
	WrongExporter  error = errors.New("Wrong exporter")
)

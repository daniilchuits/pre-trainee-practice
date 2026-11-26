package report

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

func RenderAll(records []Record, exporters []Exporter) (map[string][]byte, error)

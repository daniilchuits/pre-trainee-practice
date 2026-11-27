package report

func CommaOF(e Exporter) (rune, bool) {
	if _, ok := e.(CSVExporter); ok {
		return ',', true
	}
	return 0, false
}

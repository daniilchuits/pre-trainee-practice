package report

import "errors"

func RenderAll(records []Record, exporters []Exporter) (map[string][]byte, error) {

	recMap := make(map[string][]byte)

	var (
		foundCsv  bool
		foundJson bool
	)

	for _, ex := range exporters {

		if foundCsv && foundJson {
			return recMap, nil
		}
		name := ex.Name()
		_, ok := CommaOF(ex) // using type assertion
		// to write notes to map only one time per exporter
		if !ok {

			if foundJson {
				return nil, errors.New("Two json exporters")
			}
			csvBytes, err := ex.Export(records)
			if err != nil {
				return nil, err
			}
			recMap[name] = csvBytes
			foundJson = true
		} else {

			if foundCsv {
				return nil, errors.New("Two csv exporters")
			}
			jsonBytes, err := ex.Export(records)
			if err != nil {
				return nil, err
			}
			recMap[name] = jsonBytes
			foundCsv = true
		}
	}
	return recMap, nil
}

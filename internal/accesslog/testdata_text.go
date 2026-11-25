package accesslog

import "os"

func TestDataText() (string, error) {

	b, err := os.ReadFile(os.Getenv("TESTDATA_FILEPATH_ACCESS"))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

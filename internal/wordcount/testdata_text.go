package wordcount

import (
	"os"
)

func TestDataText() (string, error) {

	b, err := os.ReadFile(os.Getenv("TESTDATA_FILEPATH"))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

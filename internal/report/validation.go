package report

import (
	"errors"
	"strconv"
	"strings"
)

func recordValidation(record Record) (string, error) {

	switch {
	case strings.TrimSpace(record.SKU) == "":
		return "", errors.New("Empty sku")
	case record.Qty < 0:
		return "", errors.New("Quantity is less than zero")
	case record.PriceCents <= 0:
		return "", errors.New("Price is zero or lower")
	}
	return strconv.Itoa(
		record.PriceCents * record.Qty,
	), nil
}

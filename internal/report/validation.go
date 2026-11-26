package report

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func recordValidation(record Record) (string, error) {

	switch {
	case strings.TrimSpace(record.SKU) == "":
		return "", errors.New("Empty sku")
	case record.Qty < 0:
		return "", fmt.Errorf("Quantity is less than zero: %d", record.Qty)
	case record.PriceCents <= 0:
		return "", fmt.Errorf("Price is zero or lower: %d", record.PriceCents)
	}
	return strconv.Itoa(
		record.PriceCents * record.Qty,
	), nil
}

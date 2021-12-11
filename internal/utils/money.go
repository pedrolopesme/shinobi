package utils

import (
	"fmt"
	"strconv"
)

func RoundMoney(val float32) (float32, error) {
	rounded, err := strconv.ParseFloat(fmt.Sprintf("%.2f", val), 32)
	return float32(rounded), err
}

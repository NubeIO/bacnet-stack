package bacnet

import (
	"strconv"
	"strings"
)

func toInt(in string) (out int, err error) {
	f, err := strconv.ParseInt(strings.TrimRight(in, "\r\n"), 16, 32)
	if err != nil {
		return 0, err
	}
	return int(f), nil
}

func toFloat(in string) (out float64, err error) {
	out, err = strconv.ParseFloat(in, 64)
	if err != nil {
		return 0, err
	}
	return out, nil
}

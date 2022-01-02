package common

import "strconv"

func FormatFloat64(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

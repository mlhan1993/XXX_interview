package utils

import (
	"strconv"
	"strings"
)

// IntArrToStr converts an array of integer to strings joined by ","
// []int{1,2,3} -> "1,2,3"
func IntArrToStr(arr []int) string {
	var tmp []string

	for _, num := range arr {
		tmp = append(tmp, strconv.Itoa(num))
	}

	return strings.Join(tmp, ",")
}

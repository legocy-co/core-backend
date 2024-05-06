package helpers

import (
	"strconv"
	"strings"
)

func StringToIntArray(str string, delim string) []int {
	// Convert a string of integers separated by a delimiter to an array of integers
	// e.g. "1,2,3,4" -> [1, 2, 3, 4]
	var intArray []int
	for _, s := range strings.Split(str, delim) {
		intValue, err := strconv.Atoi(s)
		if err != nil {
			continue
		}

		intArray = append(intArray, intValue)
	}

	return intArray
}

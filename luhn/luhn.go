//Package luhn contains a function to validate credit card numbers
package luhn

import (
	"strconv"
	"strings"
)

//Valid will return true of the number is a valid luhn number
func Valid(input string) bool {
	input = strings.Replace(input, " ", "", -1)

	if len(input) < 2 {
		return false
	}

	if _, err := strconv.Atoi(input); err != nil {
		return false
	}

	inputAsRunes := []rune(input)

	sum := 0
	shouldDouble := len(inputAsRunes)%2 == 0

	for j := 0; j < len(inputAsRunes); j++ {
		v := int(inputAsRunes[j] - '0')

		if shouldDouble {
			v *= 2
			if v > 9 {
				v -= 9
			}
		}

		sum += v
		shouldDouble = !shouldDouble
	}

	return sum%10 == 0
}
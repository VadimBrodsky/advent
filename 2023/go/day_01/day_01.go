package day01

import (
	"strconv"
	"strings"
	"unicode"
)

func DecodeInput(input string) []string {
	codes := strings.Split(input, "\n")
	var numbers []string

	for _, code := range codes {
		number := extractNumberFromString(code)
		numbers = append(numbers, string(number))
	}

	return numbers
}

func GetCalibration(input []string) int {
	var sum int

	for _, el := range input {
		number, err := strconv.Atoi(el)

		if err == nil {
			sum += number
		}
	}

	return sum
}

func extractNumberFromString(s string) []rune {
	var digits []rune

	for _, c := range s {
		if !unicode.IsDigit(c) {
			continue
		}

		if len(digits) == 2 {
			digits[1] = c
		} else {
			digits = append(digits, c)
		}
	}

	if len(digits) == 1 {
		digits = append(digits, digits[0])
	}

	return digits
}

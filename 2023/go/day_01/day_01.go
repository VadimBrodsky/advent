package day01

import (
	"strings"
)

func DecodeInput(input string) []int {
	codes := strings.Split(input, "\n")
	decodedCodes := make([]int, len(codes)) 

	for i, c := range codes {
		decodedCodes[i] = replaceWordNumbers(c)
	}

	return decodedCodes
}

func SumCoordinates(input []int) int {
	var sum int

	for _, number := range input {
		sum += number
	}

	return sum
}

var numberWords = map[string]int{
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func replaceWordNumbers(s string) int {
	var minDigit, maxDigit int
	minIndex := len(s)
	maxIndex := 0

	for word, number := range numberWords {
		min := strings.Index(s, word)
		max := strings.LastIndex(s, word)

		if min != -1 && min <= minIndex {
			minIndex = min
			minDigit = number
		}

		if max != -1 && max >= maxIndex {
			maxIndex = max
			maxDigit = number
		}
	}

	return minDigit*10 + maxDigit
}

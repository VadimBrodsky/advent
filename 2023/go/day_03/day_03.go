package day03

import (
	"fmt"
	"regexp"
	"strings"
)

type Schematic struct {
	input string
	lines []string
}

type Part struct {
	value string
	y     int
	x1    int
	x2    int
}

type Symbol struct {
	value string
	y     int
	x     int
}

func GetParts(input string) (usedParts []int) {
	schematic := NewSchematic(input)
	parts, symbols := schematic.Parse(".")

	fmt.Printf("parts: %v\n", parts)
	fmt.Printf("symbols: %v\n", symbols)

	return
}

func Sum(a []int) (sum int) {
	for _, e := range a {
		sum += e
	}
	return
}

func NewSchematic(input string) Schematic {
	return Schematic{
		input: input,
		lines: strings.Split(input, "\n"),
	}
}

func (s Schematic) Parse(del string) (parts []Part, symbols []Symbol) {
	for y, line := range s.lines {
		symbolPattern := regexp.MustCompile(`[^.\d]+`)
		symbolsMatched := symbolPattern.FindAllString(line, -1)
		symbolsIndex := symbolPattern.FindAllStringIndex(line, -1)

		if len(symbolsIndex) != 0 && len(symbolsMatched) != 0 {
			for i, el := range symbolsIndex {
				s := Symbol{x: el[0], y: y, value: symbolsMatched[i]}
				symbols = append(symbols, s)
			}
		}

		numberPattern := regexp.MustCompile(`\d+`)
		numbersMatched := numberPattern.FindAllString(line, -1)
		numbersIndex := numberPattern.FindAllStringIndex(line, -1)

		if len(numbersIndex) != 0 && len(numbersMatched) != 0 {
			for i, el := range numbersIndex {
				p := Part{x1: el[0], x2: el[1] - 1, y: y, value: numbersMatched[i]}
				parts = append(parts, p)
			}
		}
	}
	return
}

package day03

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

type Schematic struct {
	input string
	lines []string
}

type Part struct {
	value string
	x1    int
	x2    int
	y     int
}

type Symbol struct {
	x     int
	y     int
	value string
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
	for x, line := range s.lines {
		// fmt.Printf("line %d: %s \n", i, line)

		for y, char := range line {
			// fmt.Printf("%s", string(char))

			if string(char) != del && !unicode.IsNumber(char) {
				symbol := Symbol{x: x, y: y, value: string(char)}
				symbols = append(symbols, symbol)
			}

			numberPatter := regexp.MustCompile(`\d+`)
			matches := numberPatter.FindIndex([]byte(line))

			fmt.Printf("matches: %v\n", matches)

			if unicode.IsNumber(char) {
				part := Part{value: string(char), x1: x, x2: x, y: y}
				parts = append(parts, part)
			}
		}

		fmt.Printf("\n")
	}

	return
}

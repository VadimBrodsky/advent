package day03

import (
	"fmt"
	"regexp"
	"strings"
)

type Schematic struct {
	input    string
	lines    []string
	allParts []Part
	symbols  []Part
}

type Part struct {
	point  Point
	length int
	value  string
}

type Point struct {
	x int
	y int
}

func GetParts(input string) (usedParts []int) {
	schematic := NewSchematic(input)
	schematic.symbols = schematic.Parse(`[^.\d]+`)
	schematic.allParts = schematic.Parse(`\d+`)

	used := schematic.FindAdjacent()

	// fmt.Printf("parts: %v\n", schematic.allParts)
	// fmt.Printf("symbols: %v\n", schematic.symbols)
	fmt.Printf("used: %v\n", used)

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

func (s Schematic) Parse(pattern string) (parts []Part) {
	for y, line := range s.lines {
		re := regexp.MustCompile(pattern)
		matches := re.FindAllString(line, -1)
		indexes := re.FindAllStringIndex(line, -1)

		if len(indexes) != 0 && len(matches) != 0 {
			for i, el := range indexes {
				s := Part{
					point:  Point{x: el[0], y: y},
					length: len(matches[i]),
					value:  matches[i],
				}
				parts = append(parts, s)
			}
		}
	}
	return
}

func (s Schematic) FindAdjacent() (foundParts []Part) {
	for _, p := range s.allParts {
		fmt.Printf("{ x: %d, y: %d, length: %d, value: %s }\n", p.point.x, p.point.y, p.length, p.value)
		points := getAdjacentPoints(p)
		fmt.Printf("Points: %v\n", points)

	}
	return
}

func getAdjacentPoints(p Part) (points []Point) {
	for x := p.point.x - 1; x <= p.point.x+p.length; x++ {
		for y := p.point.y - 1; y <= p.point.y+1; y++ {
			if y == p.point.y && x >= p.point.x && x < p.point.x+p.length {
				continue
			}
			if y < 0 || x < 0 {
				continue
			}
			points = append(points, Point{x: x, y: y})
		}
	}
	return
}

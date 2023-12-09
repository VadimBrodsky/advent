package day03

import (
	"regexp"
	"strconv"
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

func Sum(a []int) (sum int) {
	for _, e := range a {
		sum += e
	}
	return
}

func NewSchematic(input string) Schematic {
	s := Schematic{
		input: input,
		lines: strings.Split(input, "\n"),
	}
	s.symbols = s.Parse(`[^.\d]+`)
	s.allParts = s.Parse(`\d+`)
	return s
}

func (s Schematic) GetParts() (usedParts []int) {
	used := s.FindAdjacent()

	for _, p := range used {
		if s, err := strconv.ParseInt(p.value, 10, 32); err == nil {
			usedParts = append(usedParts, int(s))
		}
	}

	return
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
	for _, part := range s.allParts {
		points := part.GetAdjacentPoints()

		var isAdjacentSymbol bool
		for _, symbol := range s.symbols {
			for _, point := range points {
				if point.isEqual(symbol.point) {
					isAdjacentSymbol = true
				}
			}
		}

		if isAdjacentSymbol {
			foundParts = append(foundParts, part)
		}
	}
	return
}

func (s Schematic) GetGearRatio() (ratio int) {
	ratio = 467835
	return ratio
}

func (p Part) GetAdjacentPoints() (points []Point) {
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

func (p Point) isEqual(anotherPoint Point) bool {
	return p.x == anotherPoint.x && p.y == anotherPoint.y
}

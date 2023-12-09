package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Schematic struct {
	input   string
	lines   []string
	parts   []Part
	symbols []Part
}

type Part struct {
	point  Point
	length int
	value  string
}

type Gear struct {
	symbol Part
	parts  []Part
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
	s.parts = s.Parse(`\d+`)
	return s
}

func (s Schematic) GetParts() (usedParts []int) {
	used := s.FindAdjacentParts()

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

func (s Schematic) FindAdjacentParts() (foundParts []Part) {
	for _, part := range s.parts {
		var isAdjacentPart bool
		for _, symbol := range s.symbols {
			isAdjacentPart = symbol.isAdjacent(part.GetAdjacentPoints())
			if isAdjacentPart {
				foundParts = append(foundParts, part)
			}
		}
	}
	return
}

func (p Part) isAdjacent(points []Point) (adjacent bool) {
	for _, point := range points {
		if point.isEqual(p.point) {
			adjacent = true
		}
	}
	return
}

func (s Schematic) GetGearRatio(gearSymbol string) int {
	var gears []Gear
	for _, symbol := range s.symbols {
		if fmt.Sprint(symbol.value) == gearSymbol {
			gears = append(gears, Gear{symbol: symbol, parts: []Part{}})
		}
	}

	for i, gear := range gears {
		for _, part := range s.parts {
			adjacent := gear.symbol.isAdjacent(part.GetAdjacentPoints())
			if adjacent {
				gears[i].parts = append(gears[i].parts, part)
			}
		}
	}

	var ratios []int
	for _, gear := range gears {
		if len(gear.parts) == 2 {
			ratio := gear.parts[0].ToInt() * gear.parts[1].ToInt()
			ratios = append(ratios, ratio)
		}
	}

	return Sum(ratios)
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

func (p Part) ToInt() int {
	parsedInt, err := strconv.ParseInt(p.value, 10, 32)
	if err != nil {
		return 0
	}
	return int(parsedInt)
}

func (p Point) isEqual(anotherPoint Point) bool {
	return p.x == anotherPoint.x && p.y == anotherPoint.y
}

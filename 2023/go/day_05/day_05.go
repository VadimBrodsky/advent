package day05

import (
	"fmt"
	"strings"
)

type Mapping struct {
	seed        int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
	location    int
}

type Mappings []Mapping

func (m Mapping) String() string {
	return fmt.Sprintf("{seed: %d, soil: %d, fertilizer: %d, water: %d, light: %d, temperature: %d, humidity: %d, location: %d}", m.seed, m.soil, m.fertilizer, m.water, m.light, m.temperature, m.humidity, m.location)
}

func (m Mappings) String() (s string) {
	for _, mapping := range m {
		s += mapping.String() + "\n"
	}
	return
}

func NewAlmanac(input string) (m Mapping) {
  line := strings.ReplaceAll(input, "\n", " ")
	line = strings.ReplaceAll(line, " map", "")
	line = strings.ReplaceAll(line, " s", "\ns")
	line = strings.ReplaceAll(line, " f", "\nf")
	line = strings.ReplaceAll(line, " w", "\nw")
	line = strings.ReplaceAll(line, " t", "\nt")
	line = strings.ReplaceAll(line, " h", "\nh")
	line = strings.ReplaceAll(line, " l", "\nl")
	parts := strings.Split(line, "\n")

	fmt.Printf("%q", parts)
	return
}

func (m Mapping) GetMappings() (mappings Mappings) {
	return
}

func (m Mapping) GetLowestLocation() (location int) {
	return location
}

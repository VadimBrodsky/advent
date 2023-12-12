package day05

import "fmt"

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
	return
}

func (m Mapping) GetMappings() (mappings Mappings) {
	return
}

func (m Mapping) GetLowestLocation() (location int) {
	return location
}

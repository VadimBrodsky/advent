package day05

import (
	"fmt"
	"regexp"
	"strconv"
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

type Relation struct {
	from   int
	to     int
	length int
}

type Relations []Relation

type Almanac struct {
	seeds                 []int
	seedToSoil            Relations
	soilToFertilizer      Relations
	fertilizerToWater     Relations
	waterToLight          Relations
	lightToTemperature    Relations
	temperatureToHumidity Relations
	humidityToLocation    Relations
}

const (
	seeds = "seeds"
	s2s   = "seed-to-soil"
	s2f   = "soil-to-fertilizer"
	f2w   = "fertilizer-to-water"
	w2l   = "water-to-light"
	l2t   = "light-to-temperature"
	t2h   = "temperature-to-humidity"
	h2l   = "humidity-to-location"
)

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

func (a Almanac) String() (s string) {
	return fmt.Sprintf(`{
seeds: %v,
seedToSoil: %v,
soilToFertilizer: %v,
fertilizerToWater: %v,
waterToLight %v,
lightToTemperature: %v,
temperatureToHumidity: %v,
humidityToLocation: %v,
}`, a.seeds, a.seedToSoil, a.soilToFertilizer, a.fertilizerToWater, a.waterToLight, a.lightToTemperature, a.temperatureToHumidity, a.humidityToLocation)

}

func NewAlmanac(input string) (m Mapping) {
	input = replaceAll(input, map[string]string{
		"\n":   " ",
		" map": "",
		" s":   "\ns",
		" f":   "\nf",
		" w":   "\nw",
		" t":   "\nt",
		" h":   "\nh",
		" l":   "\nl",
	})
	lines := strings.Split(input, "\n")
	al := parseCategories(lines)

	fmt.Printf("\n%v", al)
	return
}

func (m Mapping) GetMappings() (mappings Mappings) {
	return
}

func (m Mapping) GetLowestLocation() (location int) {
	return location
}

func replaceAll(s string, replacements map[string]string) string {
	for old, new := range replacements {
		s = strings.ReplaceAll(s, old, new)
	}
	return s
}

func appendRelation(r Relations, values []int) Relations {
	for i := 0; i < len(values); i += 3 {
		r = append(r, Relation{
			to:     values[i],
			from:   values[i+1],
			length: values[i+2],
		})

	}
	return r
}

func getNumbers(s string) (numbers []int) {
	re := regexp.MustCompile(`\d+`)
	numberSlice := re.FindAllString(s, -1)

	for _, s := range numberSlice {
		parsedInt, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			fmt.Printf("cannot parse number %q\n", s)
		}
		numbers = append(numbers, int(parsedInt))
	}
	return
}

func parseCategories(lines []string) (al Almanac) {
	for _, line := range lines {
		category, relNumbers, _ := strings.Cut(line, ":")
		rel := getNumbers(relNumbers)

		switch category {
		case seeds:
			al.seeds = rel
		case s2s:
			al.seedToSoil = appendRelation(al.seedToSoil, rel)
		case s2f:
			al.soilToFertilizer = appendRelation(al.soilToFertilizer, rel)
		case f2w:
			al.fertilizerToWater = appendRelation(al.fertilizerToWater, rel)
		case w2l:
			al.waterToLight = appendRelation(al.waterToLight, rel)
		case l2t:
			al.lightToTemperature = appendRelation(al.lightToTemperature, rel)
		case t2h:
			al.temperatureToHumidity = appendRelation(al.temperatureToHumidity, rel)
		case h2l:
			al.humidityToLocation = appendRelation(al.humidityToLocation, rel)
		}
	}
	return
}

package day05

import (
	"fmt"
	"regexp"
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
	from   string
	to     string
	length string
}

type Relations []Relation

type Almanac struct {
	seeds                 []string
	seedToSoil            Relations
	soilToFertilizer      Relations
	fertilizerToWater     Relations
	waterToLight          Relations
	lightToTemperature    Relations
	temperatureToHumidity Relations
	humidityToLocation    Relations
}

var Categoies = []string{
	"seeds",
	"seed-to-soil",
	"soil-to-fertilizer",
	"fertilizer-to-water",
	"water-to-light",
	"light-to-temperature",
	"temperature-to-humidity",
	"humidity-to-location",
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

func (a Almanac) String() (s string) {
	return fmt.Sprintf(`{
seeds: %q,
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
	replacements := map[string]string{
		"\n":   " ",
		" map": "",
		" s":   "\ns",
		" f":   "\nf",
		" w":   "\nw",
		" t":   "\nt",
		" h":   "\nh",
		" l":   "\nl",
	}
	for old, new := range replacements {
		input = strings.ReplaceAll(input, old, new)
	}

	lines := strings.Split(input, "\n")
	var al Almanac
	re := regexp.MustCompile(`\d+`)
	for _, line := range lines {
		l := strings.Split(line, ":")

		switch {
		case strings.Contains(l[0], "seeds"):
			al.seeds = re.FindAllString(l[1], -1)
		case strings.Contains(l[0], "seed-to-soil"):
			al.seedToSoil.append(re.FindAllString(l[1], -1))
		case strings.Contains(l[0], "soil-to-fertilizer"):
			al.soilToFertilizer.append(re.FindAllString(l[1], -1))
		case strings.Contains(l[0], "fertilizer-to-water"):
			al.fertilizerToWater.append(re.FindAllString(l[1], -1))
		case strings.Contains(l[0], "water-to-light"):
			al.waterToLight.append(re.FindAllString(l[1], -1))
		case strings.Contains(l[0], "light-to-temperature"):
			al.lightToTemperature.append(re.FindAllString(l[1], -1))
		case strings.Contains(l[0], "temperature-to-humidity"):
			al.temperatureToHumidity.append(re.FindAllString(l[1], -1))
		case strings.Contains(l[0], "humidity-to-location"):
			al.humidityToLocation.append(re.FindAllString(l[1], -1))
		}

		fmt.Printf("%q", line)
	}

	fmt.Printf("\n%v", al)
	return
}

func (r *Relations) append(values []string) {
	for i := 0; i < len(values); i += 3 {
		*r = append(*r, Relation{
			to:     values[i],
			from:   values[i+1],
			length: values[i+2],
		})

	}
}

func (m Mapping) GetMappings() (mappings Mappings) {
	return
}

func (m Mapping) GetLowestLocation() (location int) {
	return location
}

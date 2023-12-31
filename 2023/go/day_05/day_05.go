package day05

import (
	"fmt"
	"math"
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

type Mappings []Mapping

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

func NewAlmanac(input string) (al Almanac) {
	input = strings.ReplaceAll(input, "\n", " ")
	input = replaceAll(input, map[string]string{
		" map": "",
		" s":   "\ns",
		" f":   "\nf",
		" w":   "\nw",
		" t":   "\nt",
		" h":   "\nh",
		" l":   "\nl",
	})
	lines := strings.Split(input, "\n")
	al = parseCategories(lines)
	return
}

func (al Almanac) GetMappings() (mappings Mappings) {
	for _, seed := range al.seeds {
		mappings = append(mappings, al.GetSeedMapping(seed))
	}
	return
}

func (al Almanac) GetSeedMapping(seed int) Mapping {
	soil := al.seedToSoil.MatchAll(seed)
	fertilizer := al.soilToFertilizer.MatchAll(soil)
	water := al.fertilizerToWater.MatchAll(fertilizer)
	light := al.waterToLight.MatchAll(water)
	temperature := al.lightToTemperature.MatchAll(light)
	humidity := al.temperatureToHumidity.MatchAll(temperature)
	location := al.humidityToLocation.MatchAll(humidity)

	return Mapping{
		seed:        seed,
		soil:        soil,
		fertilizer:  fertilizer,
		water:       water,
		light:       light,
		temperature: temperature,
		humidity:    humidity,
		location:    location,
	}
}

func (al Almanac) GetLowestLocation() (location int) {
	location = math.MaxInt

	for _, i := range al.seeds {
		mapping := al.GetSeedMapping(i)
		if mapping.location < location {
			location = mapping.location
		}
	}
	return
}

func (al Almanac) GetLowestLocationForRange() (location int) {
	location = math.MaxInt
	for i := 0; i < len(al.seeds); i += 2 {
		for j := al.seeds[i]; j < al.seeds[i]+al.seeds[i+1]; j++ {
			mapping := al.GetSeedMapping(j)
			if mapping.location < location {
				location = mapping.location
			}
		}
	}
	return
}

func parseCategories(lines []string) (al Almanac) {
	for _, line := range lines {
		category, relNumbers, _ := strings.Cut(line, ":")
		numbers := getNumbers(relNumbers)

		switch category {
		case seeds:
			al.seeds = numbers
		case s2s:
			al.seedToSoil = appendRelation(al.seedToSoil, numbers)
		case s2f:
			al.soilToFertilizer = appendRelation(al.soilToFertilizer, numbers)
		case f2w:
			al.fertilizerToWater = appendRelation(al.fertilizerToWater, numbers)
		case w2l:
			al.waterToLight = appendRelation(al.waterToLight, numbers)
		case l2t:
			al.lightToTemperature = appendRelation(al.lightToTemperature, numbers)
		case t2h:
			al.temperatureToHumidity = appendRelation(al.temperatureToHumidity, numbers)
		case h2l:
			al.humidityToLocation = appendRelation(al.humidityToLocation, numbers)
		}
	}
	return
}

func (m Mapping) String() string {
	return fmt.Sprintf("{seed: %d, soil: %d, fertilizer: %d, water: %d, light: %d, temperature: %d, humidity: %d, location: %d}", m.seed, m.soil, m.fertilizer, m.water, m.light, m.temperature, m.humidity, m.location)
}

func (m Mappings) String() (s string) {
	for _, mapping := range m {
		s += mapping.String() + "\n"
	}
	return
}

func (r Relations) MatchAll(i int) (matched int) {
	for _, relation := range r {
		matched = relation.Match(i)

		// break as soon as we get a match
		if matched != 0 && matched != i {
			break
		}
	}
	return
}

func (rel Relation) Match(input int) (matched int) {
	matched = input
	if input >= rel.from && input < rel.from+rel.length {
		diff := input - rel.from
		diff = int(math.Abs(float64(diff)))
		matched = rel.to + diff
	}
	return
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
		parsedInt, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			fmt.Printf("cannot parse number %q\n", s)
		}
		numbers = append(numbers, int(parsedInt))
	}
	return
}

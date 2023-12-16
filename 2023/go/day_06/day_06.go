package day06

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Races []Race

type Race struct {
	time     int
	distance int
}

const (
	TIME         = "Time"
	DISTANCE     = "Distance"
	ACCELERATION = 1
)

func NewBoatRaces(input string) (r Races) {
	lines := strings.Split(input, "\n")
	times, distances, err := parseLines(lines)

	if err != nil {
		panic("something broke: " + err.Error())
	}

	for i, time := range times {
		race := Race{time: time, distance: distances[i]}
		r = append(r, race)
	}

	fmt.Printf("races: %v\n", r)
	return
}

func (r Races) Wins() (wins []int) {
	wins = make([]int, len(r))

	for i, race := range r {
		for t := 0; t < race.time; t++ {
			speed := t * ACCELERATION
			timeLeft := race.time - t
			distanceTraveled := speed * timeLeft

			if distanceTraveled > race.distance {
				wins[i] += 1
			}
		}
	}

	return
}

func (r Races) WinsProduct() (p int) {
	p = 1

	wins := r.Wins()
	for _, win := range wins {
		p *= win
	}

	return
}

func (r Race) String() string {
	return fmt.Sprintf("{time: %d, distance: %d}", r.time, r.distance)
}

func parseLines(lines []string) (times, distances []int, err error) {
	for _, line := range lines {
		dataType, results, e := parseLine(line)
		if e != nil {
			return []int{}, []int{}, e
		}

		for _, result := range results {
			if dataType == TIME {
				times = append(times, result)
			}

			if dataType == DISTANCE {
				distances = append(distances, result)
			}
		}
	}
	if len(times) != len(distances) {
		err = fmt.Errorf("Mismatch in times and distances")
	}
	return
}

func parseLine(s string) (dataType string, results []int, err error) {
	dataType, values, found := strings.Cut(s, ":")
	if !found {
		err = fmt.Errorf("bad input data")
	}

	re := regexp.MustCompile(`\d+`)
	valuesSlice := re.FindAllString(values, -1)

	for _, v := range valuesSlice {
		number, e := strconv.Atoi(v)
		if e != nil {
			err = e
		}
		results = append(results, number)
	}

	return
}

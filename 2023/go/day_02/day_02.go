package day02

import (
	"fmt"
	"strconv"
	"strings"
)

type ParsedGame [][]set
type set map[string]int

const (
	RED   = "red"
	GREEN = "green"
	BLUE  = "blue"
)

func GetGames(games string, configuration map[string]int) []int {
	parsedGames := parseGames(games)
	var possibleGames []int

	for i, game := range parsedGames {

		var validGame bool = true
		for _, set := range game {
			fmt.Printf("game %d, set %v \n", i, set)
			if (configuration[RED] < set[RED]) || (configuration[BLUE] < set[BLUE]) || (configuration[GREEN] < set[GREEN]) {
				validGame = false
			}
		}

		if validGame {
			possibleGames = append(possibleGames, i+1)
		}
	}

	return possibleGames
}

func Sum(input []int) int {
	var sum int
	for _, v := range input {
		sum += v
	}
	return sum
}


func parseGames(input string) ParsedGame {
	trimmedInput := strings.TrimRight(input, "\n")
	games := strings.Split(trimmedInput, "\n")
	var parsedGames = make(ParsedGame, len(games))

	for i, game := range games {
		_, cubesString, ok := strings.Cut(game, ": ")

		if !ok {
			return parsedGames
		}

		sets := strings.Split(cubesString, "; ")
		var parsedSets []set
		for _, set := range sets {
			r, g, b := parseSet(set)
			parsedSets = append(parsedSets, map[string]int{RED: r, GREEN: g, BLUE: b})
		}
		parsedGames[i] = parsedSets
	}
	return parsedGames
}

func parseSet(set string) (r, g, b int) {
	cubes := strings.Split(set, ", ")

	for _, cube := range cubes {
		numberColourTuple := strings.Split(cube, " ")
		number := parseNumber(numberColourTuple[0])
		color := parseColour(numberColourTuple[1])

		switch color {
		case RED:
			r += number
		case GREEN:
			g += number
		case BLUE:
			b += number
		}
	}

	return
}

func parseNumber(s string) int {
	n, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0
	}
	return int(n)
}

func parseColour(s string) string {
	switch s {
	case "red":
		return RED
	case "blue":
		return BLUE
	case "green":
		return GREEN
	default:
		return ""
	}
}

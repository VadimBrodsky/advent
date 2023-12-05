package day02

import (
	"fmt"
	"strconv"
	"strings"
)

type ParsedGame []map[string]int

const (
	RED   = "red"
	GREEN = "green"
	BLUE  = "blue"
)

func GetGames(games string, configuration map[string]int) []int {
	parsedGames := parseGames(games)
	var possibleGames []int

	for _, game := range parsedGames {
		for k, v := range game {
			fmt.Printf("k: %s v: %d\n", k, v)

		}
	}

	fmt.Printf("parsedGames %v\n\n", parsedGames)
	return possibleGames
}

func Sum(input []int) int {
	return 0
}

func parseGames(input string) []map[string]int {
	games := strings.Split(input, "\n")
	var parsedGames = make(ParsedGame, len(games))

	for i, game := range games {
		_, cubesString, ok := strings.Cut(game, ": ")

		if !ok {
			return parsedGames
		}

		var gameR, gameG, gameB int

		sets := strings.Split(cubesString, "; ")
		for _, set := range sets {
			r, g, b := parseSet(set)
			gameR += r
			gameG += g
			gameB += b
		}

		parsedGames[i] = map[string]int{RED: gameR, GREEN: gameG, BLUE: gameB}
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

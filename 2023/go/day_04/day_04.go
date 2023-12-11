package day04

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	numbers        []int
	winningNumbers []int
}

type Cards []Card

type NumberOfCards []int

func newCards(s string) (cards Cards) {
	cardsStrings := strings.Split(s, "\n")
	for _, c := range cardsStrings {
		_, numbers, _ := strings.Cut(c, ": ")
		n, wn, _ := strings.Cut(numbers, "|")

		cards = append(cards, Card{
			numbers:        parseNumbers(n),
			winningNumbers: parseNumbers(wn),
		})
	}

	return
}

func (cards Cards) GetWinningNumbers() (winningNumbers [][]int) {
	for _, card := range cards {
		var winningNumbersForCard []int
		for _, wn := range card.winningNumbers {
			for _, n := range card.numbers {
				if n == wn {
					winningNumbersForCard = append(winningNumbersForCard, wn)
				}
			}
		}
		winningNumbers = append(winningNumbers, winningNumbersForCard)
	}
	return
}

func (c Cards) GetPoints() (points []int) {
	winningNumbers := c.GetWinningNumbers()

	for _, game := range winningNumbers {
		var pointsForGame int
		for i := range game {
			switch i {
			case 0:
				pointsForGame = 1
			case 1:
				pointsForGame = 2
			default:
				pointsForGame *= 2
			}
		}
		points = append(points, pointsForGame)

	}
	return
}

func (c Cards) GetScore() (score int) {
	points := c.GetPoints()
	for _, point := range points {
		score += point
	}
	return
}

func (c Cards) GetTotalScratchCards() (num NumberOfCards ){
	return
}

func (cards NumberOfCards) Sum() (s int) {
	for _, n := range cards {
		s += n
	}
	return
}

func parseNumbers(numberString string) (numbers []int) {
	re := regexp.MustCompile(`\d+`)
	numberSlice := re.FindAllString(numberString, -1)

	for _, s := range numberSlice {
		parsedInt, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			fmt.Printf("cannot parse number %q\n", s)
		}
		numbers = append(numbers, int(parsedInt))
	}

	return
}

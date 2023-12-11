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
	wins           int
	copies         int
}

type Cards []Card

type NumberOfCards []int

func newCards(s string) (cards Cards) {
	s = strings.TrimSpace(s)
	cardsStrings := strings.Split(s, "\n")

	for _, c := range cardsStrings {
		_, numbers, _ := strings.Cut(c, ": ")
		wn, n, _ := strings.Cut(numbers, "|")

		cards = append(cards, Card{
			winningNumbers: parseNumbers(wn),
			numbers:        parseNumbers(n),
			copies:         1,
		})
	}

	return
}

func (cards Cards) GetWinningNumbers() (winningNumbers [][]int) {
	for i, card := range cards {
		var winningNumbersForCard []int
		for _, wn := range card.winningNumbers {
			for _, n := range card.numbers {
				if n == wn {
					winningNumbersForCard = append(winningNumbersForCard, wn)
				}
			}
		}
		winningNumbers = append(winningNumbers, winningNumbersForCard)
		cards[i].wins = len(winningNumbersForCard)
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

func (cards Cards) GetTotalScratchCards() (num NumberOfCards) {
	winningCards := cards.GetWinningNumbers()

	for i := 0; i < len(cards); i++ {
		for copy := 0; copy < cards[i].copies; copy++ {
			for j := i + 1; j < len(winningCards[i])+i+1; j++ {
				cards[j].copies += 1
			}
		}
	}

	for _, card := range cards {
		num = append(num, card.copies)
	}
	return
}

func (cards NumberOfCards) Sum() (s int) {
	for _, n := range cards {
		s += n
	}
	return
}

func (c Card) String() string {
	return fmt.Sprintf("wn: %v, n: %v, w: %d, c: %d\n", c.winningNumbers, c.numbers, c.wins, c.copies)
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

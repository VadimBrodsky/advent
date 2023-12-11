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

func (cards Cards) GetTotalScratchCards() (num NumberOfCards) {
	winningCards := cards.GetWinningNumbers()

	for i := range cards {
		cards[i].wins += 1

		for j := i + 1; j < i+1+len(winningCards[i]); j++ {
			cards[j].wins += 1
		}
		// fmt.Printf("%v, %d\n", c, nWins)
	}

	fmt.Printf("%v\n", cards)
	return
}

func (cards NumberOfCards) Sum() (s int) {
	for _, n := range cards {
		s += n
	}
	return
}

func (c Card) String() string {
	return fmt.Sprintf("numbers: %v, winningNumbers: %v, wins: %d\n", c.numbers, c.winningNumbers, c.wins)
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

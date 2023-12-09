package day04

import (
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	numbers        []int
	winningNumbers []int
}

type Cards []Card

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

	fmt.Printf("Cards: %v\n", cards)
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
	return
}

func (c Cards) GetScore() (score int) {
	return
}

func parseNumbers(numberString string) (numbers []int) {
	numberSlice := strings.Split(numberString, " ")
	for _, s := range numberSlice {

		parsedInt, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			numbers = append(numbers, 0)
		}
		numbers = append(numbers, int(parsedInt))
	}

	return
}

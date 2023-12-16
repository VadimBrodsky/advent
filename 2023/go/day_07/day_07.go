package day07

import (
	"bufio"
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bid   int
}

type Hands []Hand

func NewCamelGame(input string) (game Hands) {
	reader := bufio.NewScanner(strings.NewReader(input))
	game, err := game.ParseHandsAndBids(reader)

	if err != nil {
		panic("something broke: " + err.Error())
	}

	fmt.Printf("game: %v\n", game)
	m := game[0].handToMap()
	fmt.Printf("mapp: %v, len: %d\n", m, len(m))
	return
}

func (h Hands) SortByRanks() (hands Hands) {
	return
}

func (h Hands) Winnings() (winnings int) {
	return
}

func (h Hands) ParseHandsAndBids(reader *bufio.Scanner) (game Hands, err error) {
	for reader.Scan() {
		line := reader.Text()
		cards, b, found := strings.Cut(line, " ")

		if !found {
			err = errors.New("Failed parsing input line")
		}

		bid, intErr := strconv.Atoi(b)
		if intErr != nil {
			err = intErr
		}

		game = append(game, Hand{cards: cards, bid: bid})
	}
	return
}

func (h Hand) String() string {
	return fmt.Sprintf("{ cards: %v, bid: %d }", h.cards, h.bid)
}

func (h Hand) handToMap() (handMap map[string]int) {
	re := regexp.MustCompile(`[AKQJT98765432]`)
	matches := re.FindAllString(h.cards, -1)
	handMap = make(map[string]int, len(matches))

	for _, match := range matches {
		handMap[match]++
	}

	return handMap
}

// where all five cards have the same label: AAAAA
func (h Hand) IsFiveOfAKind() (result bool) {
	return len(h.handToMap()) == 1
}

// where four cards have the same label and one card has a different label: AA8AA
func (h Hand) IsFourOfAKind() (result bool) {
	values := getSortedValues(h.handToMap())

	if slices.Equal(values, []int{1, 4}) {
		result = true
	}
	return
}

// where three cards have the same label, and the remaining two cards share a different label: 23332
func (h Hand) IsFullHouse() (result bool) {
	values := getSortedValues(h.handToMap())

	if slices.Equal(values, []int{2, 3}) {
		result = true
	}
	return
}

// where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
func (h Hand) IsThreeOfAKind() (result bool) {
	values := getSortedValues(h.handToMap())

	if slices.Equal(values, []int{1, 1, 3}) {
		result = true
	}
	return
}

// where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
func (h Hand) IsTwoPair() (result bool) {
	values := getSortedValues(h.handToMap())

	if slices.Equal(values, []int{1, 2, 2}) {
		result = true
	}
	return
}

// where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
func (h Hand) IsOnePair() (result bool) {
	values := getSortedValues(h.handToMap())

	if slices.Equal(values, []int{1, 1, 1, 2}) {
		result = true
	}
	return
}

// where all cards' labels are distinct: 23456
func (h Hand) IsHighCard() (result bool) {
	return len(h.handToMap()) == 5
}

func getSortedValues(m map[string]int) (values []int) {
	for _, v := range m {
		values = append(values, v)
	}

	slices.Sort(values)
	return
}

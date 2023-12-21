package day07

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Game struct {
	hands  Hands
	jokers bool
}

type Hand struct {
	cards string
	bid   int
}

type Hands []Hand

func NewCamelGame(input io.Reader, jokers bool) (game Game) {
	reader := bufio.NewScanner(input)
	game, err := game.ParseHandsAndBids(reader)

	if err != nil {
		panic("something broke: " + err.Error())
	}

	return
}

func (g Game) SortByRanks() Hands {
	sort.Sort(g.hands)
	return g.hands
}

func (g Game) Winnings() (winnings int) {
	for i, hand := range g.hands {
		winnings += (i + 1) * hand.bid
	}
	return
}

func (h Hands) Len() int {
	return len(h)
}

func (h Hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Hands) Less(i, j int) bool {
	iScore := h[i].checkRules()
	jScore := h[j].checkRules()

	if iScore == jScore {
		for k := range h[i].cards {
			iCard, jCard := h[i].cards[k], h[j].cards[k]
			if iCard == jCard {
				continue
			}

			iScore = cardToScore(iCard)
			jScore = cardToScore(jCard)

			break
		}
	}
	return iScore < jScore
}

func (g Game) ParseHandsAndBids(reader *bufio.Scanner) (game Game, err error) {
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

		game.hands = append(game.hands, Hand{cards: cards, bid: bid})
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

func (h Hand) checkRules() (result int) {
	switch {
	case h.IsFiveOfAKind():
		result = 7
	case h.IsFourOfAKind():
		result = 6
	case h.IsFullHouse():
		result = 5
	case h.IsThreeOfAKind():
		result = 4
	case h.IsTwoPair():
		result = 3
	case h.IsOnePair():
		result = 2
	case h.IsHighCard():
		result = 1
	}
	return
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

func cardToScore(card byte) (score int) {
	switch card {
	case 'A':
		score = 14
	case 'K':
		score = 13
	case 'Q':
		score = 12
	case 'J':
		score = 11
	case 'T':
		score = 10
	case '9':
		score = 9
	case '8':
		score = 8
	case '7':
		score = 7
	case '6':
		score = 6
	case '5':
		score = 5
	case '4':
		score = 4
	case '3':
		score = 3
	case '2':
		score = 2
	}
	return
}

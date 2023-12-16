package day07

import (
	"bufio"
	"errors"
	"fmt"
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

package day07

import (
	"os"
	"slices"
	"strings"
	"testing"
)

func TestDay07(t *testing.T) {
	sampleHandsAndBids := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

	t.Run("get the correct total winnings for the sample input", func(t *testing.T) {
		wantSortedByRanks := Hands{
			Hand{cards: "32T3K", bid: 765},
			Hand{cards: "KTJJT", bid: 220},
			Hand{cards: "KK677", bid: 28},
			Hand{cards: "T55J5", bid: 684},
			Hand{cards: "QQQJA", bid: 483},
		}
		wantTotalWinnings := 6440

		hands := NewCamelGame(strings.NewReader(sampleHandsAndBids), false)
		gotSortedByRanks := hands.SortByRanks()
		gotTotalWinnings := hands.Winnings()

		if !slices.Equal(wantSortedByRanks, gotSortedByRanks) {
			t.Errorf("want %v, got %v", wantSortedByRanks, gotSortedByRanks)
		}

		if wantTotalWinnings != gotTotalWinnings {
			t.Errorf("want %d, got %d", wantTotalWinnings, gotTotalWinnings)
		}
	})

	t.Run("get the correct total winnings for the full input", func(t *testing.T) {
		file, err := os.Open("input.txt")
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		wantTotalWinnings := 251058093

		hands := NewCamelGame(file, false)
		hands.SortByRanks()
		gotTotalWinnings := hands.Winnings()

		if wantTotalWinnings != gotTotalWinnings {
			t.Errorf("want %d, got %d", wantTotalWinnings, gotTotalWinnings)
		}
	})

	t.Run("get the correct total winnings with jokers for the sample input", func(t *testing.T) {
		wantSortedByRanks := Hands{
			Hand{cards: "32T3K", bid: 765},
			Hand{cards: "KK677", bid: 28},
			Hand{cards: "T55J5", bid: 684},
			Hand{cards: "QQQJA", bid: 483},
			Hand{cards: "KTJJT", bid: 220},
		}
		wantTotalWinnings := 5905

		hands := NewCamelGame(strings.NewReader(sampleHandsAndBids), true)
		gotSortedByRanks := hands.SortByRanks()
		gotTotalWinnings := hands.Winnings()

		if !slices.Equal(wantSortedByRanks, gotSortedByRanks) {
			t.Errorf("want %v, got %v", wantSortedByRanks, gotSortedByRanks)
		}

		if wantTotalWinnings != gotTotalWinnings {
			t.Errorf("want %d, got %d", wantTotalWinnings, gotTotalWinnings)
		}
	})
}

package day07

import (
	"slices"
	"testing"
)

func TestDay07(t *testing.T) {
	sampleHandsAndBids := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

	t.Run("get the correct total winnings for the sample input", func(t *testing.T) {
		wantSortedByRanks := Hands{"32T3K", "KTJJT", "KK677", "QQQJA", "T77J5"}
		wantTotalWinnings := 6440

		hands := NewCamelGame(sampleHandsAndBids)
		gotSortedByRanks := hands.SortByRanks()
		gotTotalWinnings := hands.Winnings()

		if !slices.Equal(wantSortedByRanks, gotSortedByRanks) {
			t.Errorf("got %v, want %v", wantSortedByRanks, gotSortedByRanks)
		}

		if wantTotalWinnings != gotTotalWinnings {
			t.Errorf("got %d, want %d", wantTotalWinnings, gotTotalWinnings)
		}
	})
}

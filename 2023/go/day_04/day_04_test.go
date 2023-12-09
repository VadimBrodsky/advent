package day04

import (
	"slices"
	"testing"
)

func TestDay04(t *testing.T) {

	sampleSratchcards := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
  `
	t.Run("part 1 sample", func(t *testing.T) {
		wantWinnignNumbers := [][]int{
			{48, 83, 17, 86},
			{32, 61},
			{1, 322},
			{84},
			{},
			{},
		}
		wantPoints := []int{8, 2, 2, 1, 0, 0}
		wantTotalScore := 13

		cards := newCards(sampleSratchcards)
		gotWinningNumbers := cards.GetWinningNumbers()
		gotPoints := cards.GetPoints()
		gotScore := cards.GetScore()

		if len(wantWinnignNumbers) != len(gotWinningNumbers) {
			t.Fatalf("want %d length of numbers, got %d", len(wantWinnignNumbers), len(gotWinningNumbers))
		}

		for i := range wantWinnignNumbers {
			if !slices.Equal(wantWinnignNumbers[i], gotWinningNumbers[i]) {
				t.Errorf("got %v, wanted %v", wantWinnignNumbers, gotWinningNumbers)
			}
		}

		if !slices.Equal(wantPoints, gotPoints) {
			t.Errorf("got %v, wanted %v", wantPoints, gotPoints)
		}

		if wantTotalScore != gotScore {
			t.Errorf("got %v, wanted %v", wantTotalScore, gotScore)
		}
	})
}

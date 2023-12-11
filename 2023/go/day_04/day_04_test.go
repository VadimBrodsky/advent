package day04

import (
	"os"
	"reflect"
	"slices"
	"testing"
)

func TestDay04(t *testing.T) {
	sampleScratchCards := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	t.Run("get the correct points ans cores for the sample input", func(t *testing.T) {
		wantWinningNumbers := [][]int{
			{48, 83, 17, 86},
			{32, 61},
			{1, 21},
			{84},
			{},
			{},
		}
		wantPoints := []int{8, 2, 2, 1, 0, 0}
		wantTotalScore := 13

		cards := newCards(sampleScratchCards)
		gotWinningNumbers := cards.GetWinningNumbers()
		gotPoints := cards.GetPoints()
		gotScore := cards.GetScore()

		if len(wantWinningNumbers) != len(gotWinningNumbers) {
			t.Fatalf("want %d length of numbers, got %d", len(wantWinningNumbers), len(gotWinningNumbers))
		}

		if reflect.DeepEqual(wantWinningNumbers, gotWinningNumbers) {
			t.Errorf("want %v, got %v", wantWinningNumbers, gotWinningNumbers)
		}

		if !slices.Equal(wantPoints, gotPoints) {
			t.Errorf("want %v, got %v", wantPoints, gotPoints)
		}

		if wantTotalScore != gotScore {
			t.Errorf("want %v, got %v", wantTotalScore, gotScore)
		}
	})

	t.Run("get the correct score for the full output", func(t *testing.T) {
		scratchCards, err := os.ReadFile("input.txt")
		if err != nil {
			t.Fatal(err)
		}

		wantScore := 27059
		cards := newCards(string(scratchCards))
		gotScore := cards.GetScore()

		if wantScore != gotScore {
			t.Errorf("want %v, got %v", wantScore, gotScore)
		}
	})
}

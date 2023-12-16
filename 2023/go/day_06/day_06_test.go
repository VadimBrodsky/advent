package day06

import (
	"os"
	"slices"
	"testing"
)

func TestDay06(t *testing.T) {
	sampleRaces := `Time:      7  15   30
Distance:  9  40  200`

	t.Run("get the correct value to beat the race for the sample input", func(t *testing.T) {
		wantWaysToWin := []int{4, 8, 9}
		wantWinProduct := 288

		races := NewBoatRaces(sampleRaces)
		gotWaysToWin := races.Wins()
		gotWinProduct := races.WinsProduct()

		if !slices.Equal(wantWaysToWin, gotWaysToWin) {
			t.Errorf("want %v, got %v", wantWaysToWin, gotWaysToWin)
		}

		if wantWinProduct != gotWinProduct {
			t.Errorf("want %v, got %v", wantWinProduct, gotWinProduct)
		}
	})

	t.Run("get the correct value to beat the race for the full input", func(t *testing.T) {
		fullRaceInput, err := os.ReadFile("input.txt")
		if err != nil {
			t.Fatal(err)
		}

		wantWinProduct := 275724

		races := NewBoatRaces(string(fullRaceInput))
		gotWinProduct := races.WinsProduct()

		if wantWinProduct != gotWinProduct {
			t.Errorf("want %v, got %v", wantWinProduct, gotWinProduct)
		}
	})
}

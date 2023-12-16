package day06

import (
	"os"
	"slices"
	"strings"
	"testing"
)

func TestDay06(t *testing.T) {
	sampleRaces := `Time:      7  15   30
Distance:  9  40  200`

	t.Run("get the correct value to beat the races for the sample input", func(t *testing.T) {
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

	t.Run("get the correct value to beat the races for the full input", func(t *testing.T) {
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

	t.Run("get the correct value to beat the single race for the sample input", func(t *testing.T) {
		wantWinProduct := 71503

		sampleRacesWithoutSpaces := strings.ReplaceAll(sampleRaces, " ", "")
		races := NewBoatRaces(sampleRacesWithoutSpaces)
		gotWinProduct := races.WinsProduct()

		if wantWinProduct != gotWinProduct {
			t.Errorf("want %v, got %v", wantWinProduct, gotWinProduct)
		}
	})

	t.Run("get the correct value to beat the single race for the full input", func(t *testing.T) {
		fullRaceInput, err := os.ReadFile("input.txt")
		if err != nil {
			t.Fatal(err)
		}

		wantWinProduct := 37286485

		fullRacesWithoutSpaces := strings.ReplaceAll(string(fullRaceInput), " ", "")
		races := NewBoatRaces(string(fullRacesWithoutSpaces))
		gotWinProduct := races.WinsProduct()

		if wantWinProduct != gotWinProduct {
			t.Errorf("want %v, got %v", wantWinProduct, gotWinProduct)
		}
	})
}

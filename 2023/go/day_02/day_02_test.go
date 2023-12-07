package day02

import (
	"os"
	"slices"
	"testing"
)

func TestDay02(t *testing.T) {
	sampleInput := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	bagConfiguration := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	t.Run("get correct sum of possible games for the sample input", func(t *testing.T) {
		expectedPossibleGames := []int{1, 2, 5}
		expectedSum := 8

		gotPossibleGames := GetPossibleGames(sampleInput, bagConfiguration)
		gotSumOfGames := Sum(gotPossibleGames)

		if !slices.Equal(expectedPossibleGames, gotPossibleGames) {
			t.Errorf("expected %v, got %v", expectedPossibleGames, gotPossibleGames)
		}

		if expectedSum != gotSumOfGames {
			t.Errorf("expected %d, got %d", expectedSum, gotSumOfGames)
		}
	})

	t.Run("get the sum of possible games for the full input", func(t *testing.T) {
		games, err := os.ReadFile("input.txt")
		if err != nil {
			t.Fatal(err)
		}

		want := 2683
		possibleGames := GetPossibleGames(string(games), bagConfiguration)
		got := Sum(possibleGames)

		if want != got {
			t.Errorf("expected sum of %d, got %d", want, got)
		}
	})

	t.Run("get the power of possible games for the sample input", func(t *testing.T) {
		expectedPowerSet := []int{48, 12, 1560, 630, 36}
		expectedSum := 2286

		gotPowerSet := GetPowerSet(sampleInput)
		gotPowerSum := Sum(gotPowerSet)

		if !slices.Equal(expectedPowerSet, gotPowerSet) {
			t.Errorf("expected %v, got %v", expectedPowerSet, gotPowerSet)
		}

		if expectedSum != gotPowerSum {
			t.Errorf("expected %d, got %d", expectedSum, gotPowerSum)
		}
	})

	t.Run("get the sum of power of possible games for the full input", func(t *testing.T) {
		games, err := os.ReadFile("input.txt")
		if err != nil {
			t.Fatal(err)
		}

		want := 49710
		powerSet := GetPowerSet(string(games))
		got := Sum(powerSet)

		if want != got {
			t.Errorf("expected sum of %d, got %d", want, got)
		}
	})
}

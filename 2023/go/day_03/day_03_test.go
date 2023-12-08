package day03

import (
	"slices"
	"testing"
)

func TestDay03(t *testing.T) {
	sampleSchematic := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	t.Run("get the correct sum for engine part numbers for the sample output", func(t *testing.T) {
		expectedPartNumbers := []int{467, 35, 633, 617, 592, 755, 664, 598}
		expectedSum := 4361

		gotPartNumbers := GetParts(sampleSchematic)
		gotSum := Sum(gotPartNumbers)

		if !slices.Equal(expectedPartNumbers, gotPartNumbers) {
			t.Errorf("expected %v, got %v", expectedPartNumbers, gotPartNumbers)
		}

		if expectedSum != gotSum {
			t.Errorf("expected sum %d, got %d", expectedSum, gotSum)
		}
	})

}

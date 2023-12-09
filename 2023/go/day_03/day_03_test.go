package day03

import (
	"os"
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

	t.Run("get the correct sum for engine part numbers for the sample input", func(t *testing.T) {
		expectedPartNumbers := []int{467, 35, 633, 617, 592, 755, 664, 598}
		expectedSum := 4361

		s := NewSchematic(sampleSchematic)
		gotPartNumbers := s.GetParts()
		gotSum := Sum(gotPartNumbers)

		if !slices.Equal(expectedPartNumbers, gotPartNumbers) {
			t.Errorf("expected %v, got %v", expectedPartNumbers, gotPartNumbers)
		}

		if expectedSum != gotSum {
			t.Errorf("expected sum %d, got %d", expectedSum, gotSum)
		}
	})

	t.Run("get the correct sum for engine part numbers for the full input", func(t *testing.T) {
		schematic, err := os.ReadFile("input.txt")
		if err != nil {
			t.Fatal(err)
		}

		expectedSum := 522726

		s := NewSchematic(string(schematic))
		gotPartNumbers := s.GetParts()
		gotSum := Sum(gotPartNumbers)

		if expectedSum != gotSum {
			t.Errorf("expected sum %d, got %d", expectedSum, gotSum)
		}
	})

	t.Run("get the correct gear ration for the sample input", func(t *testing.T) {
		expectedGearRatio := 467835

		s := NewSchematic(sampleSchematic)
		gotGearRatio := s.GetGearRatio()

		if expectedGearRatio != gotGearRatio {
			t.Errorf("expected sum %d, got %d", expectedGearRatio, gotGearRatio)
		}
	})
}

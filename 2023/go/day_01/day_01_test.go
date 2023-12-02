package day01

import (
	"os"
	"slices"
	"testing"
)

func TestDay01(t *testing.T) {
	t.Run("get correct calibration values for the sample input", func(t *testing.T) {
		input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

		desiredDecodedInput := []string{"12", "38", "15", "77"}
		want := 142

		decodedInput := DecodeInput(input)
		got := GetCalibration(decodedInput)

		if !slices.Equal(desiredDecodedInput, decodedInput) {
			t.Errorf("expected decoded input of %v, got %v", desiredDecodedInput, decodedInput)
		}

		if want != got {
			t.Errorf("expected sum of %d, got %d", want, got)
		}
	})

	t.Run("get the correct calibration value for part 2", func(t *testing.T) {
		input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

		desiredDecodedInput := []string{"29", "83", "13", "24", "42", "14", "76"}
		want := 281

		decodedInput := DecodeInput(input)
		got := GetCalibration(decodedInput)

		if !slices.Equal(desiredDecodedInput, decodedInput) {
			t.Errorf("expected decoded input of %v, got %v", desiredDecodedInput, decodedInput)
		}

		if want != got {
			t.Errorf("expected sum of %d, got %d", want, got)
		}
	})

	t.Run("get the correct calibration values for the full input", func(t *testing.T) {
		data, err := os.ReadFile("input.txt")
		if err != nil {
			t.Fatal(err)
		}

		// want := 53921
		want := 54665 
		decodedInput := DecodeInput(string(data))
		got := GetCalibration(decodedInput)

		if want != got {
			t.Errorf("expected sum of %d, got %d", want, got)
		}
	})
}

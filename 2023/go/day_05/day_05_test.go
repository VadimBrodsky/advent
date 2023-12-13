package day05

import (
	"os"
	"testing"
)

func TestDay05(t *testing.T) {
	sampleAlmanac := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

	t.Run("should return the correct locations for the sample almanac input", func(t *testing.T) {
		t.Skip("")
		wantMapping := Mappings{
			Mapping{seed: 79, soil: 81, fertilizer: 81, water: 81, light: 74, temperature: 78, humidity: 78, location: 82},
			Mapping{seed: 14, soil: 14, fertilizer: 53, water: 49, light: 42, temperature: 42, humidity: 43, location: 43},
			Mapping{seed: 55, soil: 57, fertilizer: 57, water: 53, light: 46, temperature: 82, humidity: 82, location: 86},
			Mapping{seed: 13, soil: 13, fertilizer: 52, water: 41, light: 34, temperature: 34, humidity: 35, location: 35},
		}
		wantLowestLocation := 35

		almanac := NewAlmanac(sampleAlmanac, false)
		gotMapping := almanac.GetMappings(false)
		gotLowestLocation := almanac.GetLowestLocation(gotMapping)

		if wantMapping.String() != gotMapping.String() {
			t.Errorf("expected \n%v, got \n%v", wantMapping, gotMapping)
		}

		if wantLowestLocation != gotLowestLocation {
			t.Errorf("expected %d, got %d", wantLowestLocation, gotLowestLocation)
		}
	})

	t.Run("should return the correct locations for the full almanac input", func(t *testing.T) {
		t.Skip("")
		fullAlmanac, err := os.ReadFile("input.txt")
		if err != nil {
			t.Fatal(err)
		}

		wantLowestLocation := 323142486
		almanac := NewAlmanac(string(fullAlmanac), false)
		gotMapping := almanac.GetMappings(false)
		gotLowestLocation := almanac.GetLowestLocation(gotMapping)

		if wantLowestLocation != gotLowestLocation {
			t.Errorf("expected %d, got %d", wantLowestLocation, gotLowestLocation)
		}
	})

	t.Run("should return the correct location for a range of seed of the sample almanac input", func(t *testing.T) {
		t.Skip("")
		wantLowestLocation := 46

		almanac := NewAlmanac(sampleAlmanac, true)
		gotMapping := almanac.GetMappings(true)
		gotLowestLocation := almanac.GetLowestLocation(gotMapping)

		if wantLowestLocation != gotLowestLocation {
			t.Errorf("expected %d, got %d", wantLowestLocation, gotLowestLocation)
		}
	})

	t.Run("should return the correct location for a range of seeds for the full almanac input", func(t *testing.T) {
		t.Skip("")
		fullAlmanac, err := os.ReadFile("input.txt")
		if err != nil {
			t.Fatal(err)
		}

		wantLowestLocation := 0
		almanac := NewAlmanac(string(fullAlmanac), true)
		gotMapping := almanac.GetMappings(true)
		gotLowestLocation := almanac.GetLowestLocation(gotMapping)

		if wantLowestLocation != gotLowestLocation {
			t.Errorf("expected %d, got %d", wantLowestLocation, gotLowestLocation)
		}
	})
}

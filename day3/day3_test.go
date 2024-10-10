package day3_test

import (
	"testing"

	"github.com/brainboxweb/advent-2023/day3"
	"github.com/brainboxweb/advent-2023/helpers"
	"github.com/stretchr/testify/assert"
)

const dataPath = "../data/"

// 467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..
func TestDay3(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "day3_test.txt",
			4361,
		},
		{
			dataPath + "day3.txt",
			526404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day3.Day3(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDay3_2(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "day3_test.txt",
			467835,
		},
		{
			dataPath + "day3.txt",
			84399773,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day3.Day3_2(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

package day5_test

import (
	"testing"

	"github.com/brainboxweb/advent-2023/day5"
	"github.com/brainboxweb/advent-2023/helpers"
	"github.com/stretchr/testify/assert"
)

const dataPath = "../data/"

func TestDay5(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "day5_test.txt",
			35,
		},
		{
			dataPath + "day5.txt",
			282277027,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day5.Part1(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDay5_2(t *testing.T) { // NB - needs ~35s to run
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "day5_test.txt",
			46,
		},
		{
			dataPath + "day5.txt",
			11554135,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day5.Part2(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

package day2_test

import (
	"testing"

	"github.com/brainboxweb/advent-2023/day2"
	"github.com/brainboxweb/advent-2023/helpers"
	"github.com/stretchr/testify/assert"
)

const dataPath = "../data/"

func TestPart1(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "day2_test.txt",
			8,
		},
		{
			dataPath + "day2.txt",
			2810,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day2.Part1(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "day2_test.txt",
			2286,
		},
		{
			dataPath + "day2.txt",
			69110,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day2.Part2(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

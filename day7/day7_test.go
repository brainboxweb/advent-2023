package day7_test

import (
	"testing"

	"github.com/brainboxweb/advent-2023/day7"
	"github.com/brainboxweb/advent-2023/helpers"
	"github.com/stretchr/testify/assert"
)

const dataPath = "../data/"

func TestDay7(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "day7_test.txt",
			6440,
		},
		{
			dataPath + "day7.txt",
			250254244,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day7.Day7(dataSet, false)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDay7_Part2(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "day7_test.txt",
			5905,
		},
		{
			dataPath + "day7.txt",
			250087440,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day7.Day7(dataSet, true)
			assert.Equal(t, tt.expected, result)
		})
	}
}

package day4_test

import (
	"testing"

	"github.com/brainboxweb/advent-2023/day4"
	"github.com/brainboxweb/advent-2023/helpers"
	"github.com/stretchr/testify/assert"
)

const dataPath = "../data/"

func TestDay4(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "day4_test.txt",
			13,
		},
		{
			dataPath + "day4.txt",
			27845,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day4.Part1(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDay4_2(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "day4_test.txt",
			30,
		},
		{
			dataPath + "day4.txt",
			9496801,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day4.Part2(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

package day8_test

import (
	"testing"

	"github.com/brainboxweb/advent-2023/day8"
	"github.com/brainboxweb/advent-2023/helpers"
	"github.com/stretchr/testify/assert"
)

const dataPath = "../data/"

func TestPart1(t *testing.T) {
	tests := []struct {
		data     []string
		dataFile string
		expected int
	}{
		{
			[]string{},
			dataPath + "day8_test.txt",
			2,
		},
		{
			[]string{},
			dataPath + "day8_test_2.txt",
			6,
		},
		{
			[]string{},
			dataPath + "day8.txt",
			19241,
		},
	}
	for _, tt := range tests {
		dataset := []string{}
		t.Run(tt.dataFile, func(t *testing.T) {
			if len(tt.data) == 0 {
				dataset = helpers.GetDataString(tt.dataFile)
			} else {
				dataset = tt.data
			}
			result := day8.Part1(dataset)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		data     []string
		dataFile string
		expected int
	}{
		{
			[]string{},
			dataPath + "day8_test_Part2.txt",
			6,
		},
		{
			[]string{},
			dataPath + "day8.txt",
			9606140307013,
		},
	}
	for _, tt := range tests {
		dataset := []string{}
		t.Run(tt.dataFile, func(t *testing.T) {
			if len(tt.data) == 0 {
				dataset = helpers.GetDataString(tt.dataFile)
			} else {
				dataset = tt.data
			}
			result := day8.Part2(dataset)
			assert.Equal(t, tt.expected, result)
		})
	}
}

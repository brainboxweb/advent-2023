package day9_test

import (
	"testing"

	"github.com/brainboxweb/advent-2023/day9"
	"github.com/brainboxweb/advent-2023/helpers"
	"github.com/stretchr/testify/assert"
)

const dataPath = "../data/"

func TestDay9_Pre(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "day9_test1.txt",
			18,
		},
		{
			dataPath + "day9_test2.txt",
			28,
		},
		{
			dataPath + "day9_test3.txt",
			68,
		},
		{
			dataPath + "day9_test4.txt",
			-330,
		},
		{
			dataPath + "day9_test5.txt",
			7343,
		},
		{
			dataPath + "day9_test6.txt",
			-1199863,
		},
		{
			dataPath + "day9_test7.txt",
			77288,
		},
	}
	for _, tt := range tests {
		t.Run("pre-test", func(t *testing.T) {
			dataset := helpers.GetDataString(tt.dataFile)
			result := day9.Part11(dataset)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDay9(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "day9_test.txt",
			114,
		},
		{
			dataPath + "day9.txt",
			1980437560,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataset := helpers.GetDataString(tt.dataFile)
			result := day9.Part11(dataset)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDay9_Part2(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			dataPath + "day9_test.txt",
			2,
		},
		{
			dataPath + "day9.txt",
			977,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataset := helpers.GetDataString(tt.dataFile)
			result := day9.Part2(dataset)
			assert.Equal(t, tt.expected, result)
		})
	}
}

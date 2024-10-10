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
		data     []string
		expected int
	}{
		{
			[]string{"0 3 6 9 12 15"},
			18,
		},
		{
			[]string{"1 3 6 10 15 21"},
			28,
		},
		{
			[]string{"10 13 16 21 30 45"},
			68,
		},
		{
			[]string{"6 0 -7 -15 -24 -34 -45 -57 -70 -84 -99 -115 -132 -150 -169 -189 -210 -232 -255 -279 -304"},
			-330,
		},
		{
			[]string{"7 23 47 83 135 207 303 427 583 775 1007 1283 1607 1983 2415 2907 3463 4087 4783 5555 6407"},
			7343,
		},
		{
			[]string{"6 30 63 111 201 393 794 1572 2962 5248 8696 13404 19027 24329 26511 20265 -3490 -59222 -168763 -363635 -687675"},
			-1199863,
		},
		{
			[]string{"15 13 15 38 108 260 538 995 1693 2703 4105 5988 8450 11598 15548 20425 26363 33505 42003 52018 63720"},
			77288,
		},
	}
	for _, tt := range tests {
		t.Run("pre-test", func(t *testing.T) {
			result := day9.Day9(tt.data, true)
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
			result := day9.Day9(dataset, true)
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
			result := day9.Day9(dataset, false)
			assert.Equal(t, tt.expected, result)
		})
	}
}

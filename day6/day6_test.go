package day6_test

import (
	"testing"

	"github.com/brainboxweb/advent-2023/day6"
	"github.com/stretchr/testify/assert"
)

func TestDay6(t *testing.T) {
	tests := []struct {
		races    []day6.Race
		expected int
	}{
		{
			races: []day6.Race{
				{
					Time: 7, Record: 9,
				},
				{
					Time: 15, Record: 40,
				},
				{
					Time: 30, Record: 200,
				},
			},
			expected: 288,
		},
		{
			races: []day6.Race{
				{
					Time: 46, Record: 214,
				},
				{
					Time: 80, Record: 1177,
				},
				{
					Time: 78, Record: 1402,
				},
				{
					Time: 66, Record: 1024,
				},
			},
			expected: 512295,
		},
		{
			races: []day6.Race{
				{
					Time: 71530, Record: 940200,
				},
			},
			expected: 71503,
		},
		{
			races: []day6.Race{
				{
					Time: 46807866, Record: 214117714021024,
				},
			},
			expected: 36530883,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			result := day6.Day6(tt.races)
			assert.Equal(t, tt.expected, result)
		})
	}
}

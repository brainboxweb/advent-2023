package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGame(t *testing.T) {
	tests := []struct {
		data     string
		expected Game
	}{
		{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			Game{
				index: "Game 1",
				turns: []Turn{
					{
						cubes: map[string]int{
							"blue": 3,
							"red":  4,
						},
					},
					{
						cubes: map[string]int{
							"red":   1,
							"green": 2,
							"blue":  6,
						},
					},
					{
						cubes: map[string]int{
							"green": 2,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.data, func(t *testing.T) {
			// dataSet := helpers.GetDataString(tt.dataFile)
			g := parseGame(tt.data)
			assert.Equal(t, tt.expected, g)
		})
	}
}

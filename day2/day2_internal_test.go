package day2

import (
	"testing"

	// "github.com/brainboxweb/advent-2022/helpers"

	"github.com/stretchr/testify/assert"
)

// const dataPath = "../data/"

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
							"red": 4, 
						},
					},
					{
						cubes: map[string]int{					
							"red": 1,
							"green": 2,
							"blue": 6, 
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
		// {
		// 	dataPath + "day2.txt",
		// 	12458,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.data, func(t *testing.T) {
			// dataSet := helpers.GetDataString(tt.dataFile)
			g := parseGame(tt.data)
			assert.Equal(t, tt.expected, g)
		})
	}
}

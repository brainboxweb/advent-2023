package game_test

import (
	"testing"

	"github.com/brainboxweb/advent-2023/day2/game"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		gameID   int
		turnData []map[string]int
		expected []int // colourCounts
	}{
		{
			1,
			[]map[string]int{
				{
					"blue": 3,
					"red":  4,
				},
				{
					"red":   1,
					"green": 2,
					"blue":  6,
				},
				{
					"green": 2,
				},
			},
			[]int{4, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			g := game.New(1)
			for _, turnData := range tt.turnData {
				g.AddTurn(turnData)
			}
			colorCounts := g.GetColorCounts("red")
			assert.Equal(t, tt.expected, colorCounts)
		})
	}
}

package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		expected Game
	}{
		{
			Game{
				index: 1,
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
		t.Run("test", func(t *testing.T) {
			g := New(1)
			turn1 := map[string]int{
				"blue": 3,
				"red":  4,
			}
			turn2 := map[string]int{
				"red":   1,
				"green": 2,
				"blue":  6,
			}
			turn3 := map[string]int{
				"green": 2,
			}
			g.AddTurn(turn1)
			g.AddTurn(turn2)
			g.AddTurn(turn3)

			assert.Equal(t, tt.expected, *g)

			assert.Equal(t, []int{4, 1, 0}, g.GetColorCounts("red"))
			assert.Equal(t, []int{3, 6, 0}, g.GetColorCounts("blue"))
			assert.Equal(t, []int{0, 2, 2}, g.GetColorCounts("green"))

			contents := map[string]int{
				"green": 2,
			}
			assert.False(t, g.IsPossible(contents))

			contents = map[string]int{
				"red":   5,
				"blue":  9,
				"green": 4,
			}
			assert.True(t, g.IsPossible(contents))

			contents = map[string]int{
				"green": 10,
				"blue":  10,
				"red":   10,
			}
			assert.True(t, g.IsPossible(contents))
		})
	}
}

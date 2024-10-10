package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRankIt(t *testing.T) {
	tests := []struct {
		hand     Hand
		expected int
	}{
		{
			Hand{
				Cards: "AAAAA",
				Bid:   0,
				Rank:  0,
			},
			1,
		},
		{
			Hand{
				Cards: "AA8AA",
				Bid:   0,
				Rank:  0,
			},
			2,
		},
		{
			Hand{
				Cards: "23332",
				Bid:   0,
				Rank:  0,
			},
			3,
		},
		{
			Hand{
				Cards: "TTT98",
				Bid:   0,
				Rank:  0,
			},
			4,
		},
		{
			Hand{
				Cards: "23432", // 2 pair
				Bid:   0,
				Rank:  0,
			},
			5,
		},
		{
			Hand{
				Cards: "A23A4", // 1 pair
				Bid:   0,
				Rank:  0,
			},
			6,
		},
		{
			Hand{
				Cards: "23456", // 1 pair
				Bid:   0,
				Rank:  0,
			},
			7,
		},
	}
	for _, tt := range tests {
		t.Run("gary", func(t *testing.T) { // searhc fo "gary" entries"!!!
			tt.hand.rank()
			assert.Equal(t, tt.expected, tt.hand.Rank)
		})
	}

}

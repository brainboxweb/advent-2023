package day1_test

import (
	"testing"

	"github.com/brainboxweb/advent-2023/day1"
	"github.com/brainboxweb/advent-2023/helpers"

	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"../data/day1_test.txt",
			142,
		},
		{
			"../data/day1.txt",
			55029,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day1.Day1(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDay1a(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"../data/day1a_test.txt",
			281, // 281 SHOULD work
		},
		{
			"../data/day1.txt",
			55686,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day1.Day1a(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDoCalcDay1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"1abc2",
			12,
		},
		{
			"pqr3stu8vwx",
			38,
		},
		{
			"a1b2c3d4e5f",
			15,
		},
		{
			"treb7uchet",
			77,
		},
		{
			"abc11abczzzz",
			11,
		},
		{
			"abc99abczzzz",
			99,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			// dataSet := helpers.GetData(tt.dataFile)
			result := day1.DoCalcDay1(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDoCalcDay1a(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"1asfddf2",
			12,
		},
		{
			"one111three",
			13,
		},
		{
			"2andseven",
			27,
		},
		{
			"two1nine",
			29,
		},
		{
			"eightwothree",
			83,
		},
		{
			"abcone2threexyz",
			13,
		},
		{
			"xtwone3four",
			24,
		},
		{
			"4nineeightseven2",
			42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			// dataSet := helpers.GetData(tt.dataFile)
			result := day1.DoCalcDay1a(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

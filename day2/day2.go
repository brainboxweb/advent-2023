package day2

import (
	"strconv"
	"strings"

	"github.com/brainboxweb/advent-2023/day2/game"
)

func Part1(data []string) int {
	games := []*game.Game{}
	for _, line := range data {
		index, turnsData := parseGameData(line)
		g := game.New(index)
		for _, turnData := range turnsData {
			g.AddTurn(turnData)
		}
		games = append(games, g)
	}

	gameIDcount := 0
	for _, g := range games {
		// 12 red cubes, 13 green cubes, and 14 blue cubes.
		gameOK := g.IsPossible(
			map[string]int{
				"red":   12,
				"green": 13,
				"blue":  14,
			},
		)
		if gameOK {
			gameIDcount += g.Index()
		}
	}

	return gameIDcount
}

//revive:disable:cognitive-complexity
func Part2(data []string) int {
	games := []*game.Game{}
	for _, line := range data {
		index, turnsData := parseGameData(line)
		g := game.New(index)
		for _, turnData := range turnsData {
			g.AddTurn(turnData)
		}
		games = append(games, g)
	}

	sumPowers := 0
	for _, g := range games {
		maxRed := 0
		maxGreen := 0
		maxBlue := 0

		reds := g.GetColorCounts("red")
		for _, r := range reds {
			if r > maxRed {
				maxRed = r
			}
		}
		greens := g.GetColorCounts("green")
		for _, g := range greens {
			if g > maxGreen {
				maxGreen = g
			}
		}
		blues := g.GetColorCounts("blue")
		for _, b := range blues {
			if b > maxBlue {
				maxBlue = b
			}
		}
		power := maxRed * maxGreen * maxBlue
		sumPowers += power
	}
	return sumPowers
}

//revive:enable:cognitive-complexity

func parseGameData(str string) (int, []map[string]int) {
	gameIndex := 0
	turnsData := []map[string]int{}

	parts := strings.Split(str, ": ")
	gameIndexPart := parts[0]
	indexParts := strings.Split(gameIndexPart, " ")

	gameIndex, err := strconv.Atoi(indexParts[1]) // handle error
	if err != nil {
		panic("not expected")
	}

	turns := strings.Split(parts[1], "; ")
	for _, turn := range turns {
		turnData := make(map[string]int)
		pps := strings.Split(turn, ", ")
		for _, cube := range pps {
			bits := strings.Split(cube, " ")
			count, _ := strconv.Atoi(bits[0])
			turnData[bits[1]] = count
		}
		turnsData = append(turnsData, turnData)
	}
	return gameIndex, turnsData
}

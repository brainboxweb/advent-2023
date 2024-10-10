package day2

import (
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func Day2(data []string) int {
	games := []Game{}
	for _, line := range data {
		g := parseGame(line)
		games = append(games, g)
	}
	//only 12 red cubes, 13 green cubes, and 14 blue cubes?
	gameIDcount := 0
	for i, g := range games {
		gameOK := true
		reds := g.GetColorCounts("red")
		for _, r := range reds {
			spew.Dump("reds: ", r)
			if r > 12 {
				gameOK = false
			}
		}
		greens := g.GetColorCounts("green")
		for _, g := range greens {
			if g > 13 {
				gameOK = false
			}
		}
		blues := g.GetColorCounts("blue")
		for _, b := range blues {
			if b > 14 {
				gameOK = false
			}
		}
		if gameOK {
			spew.Dump("good game: ", i+1)
			gameIDcount += i + 1
		}
	}

	return gameIDcount
}

func Day2_2(data []string) int {
	games := []Game{}
	for _, line := range data {
		g := parseGame(line)
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

func parseGame(str string) Game {
	parts := strings.Split(str, ": ")
	gameIndex := parts[0]
	g := Game{index: gameIndex}
	turns := strings.Split(parts[1], "; ")
	for _, turn := range turns {
		t := NewTurn()
		pps := strings.Split(turn, ", ")
		for _, cube := range pps {
			bits := strings.Split(cube, " ")
			count, _ := strconv.Atoi(bits[0])
			t.addCubes(bits[1], count)
		}
		g.turns = append(g.turns, t)
	}
	return g
}

type Turn struct {
	cubes map[string](int)
}

type Game struct {
	index string
	turns []Turn
}

func NewTurn() Turn {
	t := Turn{}
	t.cubes = make(map[string]int)
	return t
}

func (t *Turn) addCubes(color string, count int) {
	t.cubes[color] = t.cubes[color] + count
}

func (g *Game) GetColorCounts(color string) []int {

	// Hmm... think they
	counts := []int{}
	for _, t := range g.turns {
		counts = append(counts, t.cubes[color])
	}
	return counts
}

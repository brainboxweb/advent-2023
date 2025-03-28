package day4

import (
	"strconv"
	"strings"
)

func Part1(data []string) int {
	ret := 0
	for _, item := range data {
		winners, mine := parseCard(item)
		ret += getScore(winners, mine)
	}

	return ret
}

func Part2(data []string) int {
	ret := 0
	winList := []int{}
	for _, item := range data {
		winners, mine := parseCard(item)
		score := getMatchCount(winners, mine)
		winList = append(winList, score)
	}
	thingyList := make(map[int]int)    // stores the multipliers
	for i := 1; i < len(data)+1; i++ { // fill the map
		thingyList[i] = 1
	}
	for cardIndex, win := range winList {
		cardIndex++ // 1 based
		for i := 1; i < win+1; i++ {
			newIndex := cardIndex + i
			toAdd := thingyList[cardIndex] // The COUNT of the CURRENT card
			thingyList[newIndex] += toAdd
		}
	}
	for _, thingy := range thingyList {
		ret += thingy
	}

	return ret
}

func parseCard(input string) (winners, mine []int) {
	parts := strings.Split(input, ":")
	divided := strings.Split(parts[1], "|")
	winners = extractCards(divided[0])
	mine = extractCards(divided[1])

	return winners, mine
}

func getMatchCount(winners, mine []int) int {
	matchCount := 0
	for _, win := range winners {
		for _, mi := range mine {
			if win == mi {
				matchCount++
			}
		}
	}
	if matchCount == 0 {
		return 0
	}

	return matchCount
}

func getScore(winners, mine []int) int {
	matchCount := getMatchCount(winners, mine)
	if matchCount == 0 {
		return 0
	}
	ret := 1
	for range matchCount - 1 {
		ret *= 2
	}

	return ret
}

func extractCards(input string) []int {
	ret := []int{}
	cards := strings.Split(input, " ")
	for _, card := range cards {
		if card == "" {
			continue
		}
		cardNo, _ := strconv.Atoi(card)
		ret = append(ret, cardNo)
	}

	return ret
}

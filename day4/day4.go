package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func Day4(data []string) int {
	ret := 0
	for _, item := range data {
		winners, mine := parseCard(item)
		ret += getScore(winners, mine)
	}
	return ret
}

func Day4_2(data []string) int {
	ret := 0

	winList := []int{}
	for _, item := range data {
		winners, mine := parseCard(item)

		score := getMatchCount(winners, mine)
		winList = append(winList, score)
		fmt.Println(score)
	}

	spew.Dump(winList)

	thingyList := make(map[int]int) // stores the multipliers

	for i := 1; i < len(data)+1; i++ { // fill the map
		thingyList[i] = 1
	}

	spew.Dump(thingyList) // stores the collected carlds

	for cardIndex, win := range winList {
		// spew.Dump("THINGY", thingyList)
		
		cardIndex += 1 // 1 based

		fmt.Println("=====================\nthe card, winm " , cardIndex, win)

		// fmt.Println("---card. win", cardIndex, win )
		// thingyList[cardIndex] += 1 // gets one anyway
		// fmt.Println("---adding the DEFAULT", cardIndex)
		// if win == 0 {
		// 	continue
		// }

		// for each win, add a card... 	

		for i := 1; i < win+1; i++ {
			newIndex := cardIndex + i
			toAdd := thingyList[cardIndex] // The COUNT of the CURRENT card
			fmt.Println("\n-- adding ", toAdd, " to " , newIndex)
			thingyList[newIndex] += toAdd
		}
	}

	spew.Dump(thingyList)

	for _, thingy := range thingyList{
		ret+= thingy
	}

	return ret
}

func parseCard(input string) ([]int, []int) {
	// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	parts := strings.Split(input, ":")
	// cardName := parts[0]

	divided := strings.Split(parts[1], "|")

	// winners := extractCards[divided[0]]

	winners := extractCards(divided[0])
	mine := extractCards(divided[1])

	return winners, mine
	// theScore := getScore(winners, mine)

	// // winners := strings.Split(winnersString, " ")

	// // mine := strings.Split(mineString, " ")

	// spew.Dump("\n==\n", cardName, winners, mine)

	// fmt.Println("the score", theScore)

	// return theScore
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

	fmt.Println("MAtch count: ", matchCount)
	return matchCount

}

func getScore(winners, mine []int) int {
	// matchCount := 0
	// for _, win := range winners {
	// 	for _, mi := range mine {
	// 		if win == mi {
	// 			matchCount++
	// 		}
	// 	}
	// }
	matchCount := getMatchCount(winners, mine)
	if matchCount == 0 {
		return 0
	}

	fmt.Println("MAtch count: ", matchCount)

	ret := 1
	for i := 0; i < matchCount-1; i++ {

		ret = ret * 2
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

// func getParts(input string) ([]int, []int) {
// 	parts := strings.Split(input, ",")

// 	aRange := strings.Split(parts[0], "-")
// 	bRange := strings.Split(parts[1], "-")

// 	aL, _ := strconv.Atoi(aRange[0])
// 	aH, _ := strconv.Atoi(aRange[1])

// 	bL, _ := strconv.Atoi(bRange[0])
// 	bH, _ := strconv.Atoi(bRange[1])

// 	a := []int{aL, aH}
// 	b := []int{bL, bH}

// 	return a, b
// }

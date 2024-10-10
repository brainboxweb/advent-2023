package day7

import (
	"sort"
	"strconv"
	"strings"
)

var rankingSystem = map[string]int{
	"A": 1,
	"K": 2,
	"Q": 3,
	"J": 4,
	"T": 5,
	"9": 6,
	"8": 7,
	"7": 8,
	"6": 9,
	"5": 10,
	"4": 11,
	"3": 12,
	"2": 13,
}

var rankingSystemWithJokers = map[string]int{
	"A": 1,
	"K": 2,
	"Q": 3,
	"T": 4,
	"9": 5,
	"8": 6,
	"7": 7,
	"6": 8,
	"5": 9,
	"4": 10,
	"3": 11,
	"2": 12,
	"J": 13,
}

func Day7(data []string, useJokers bool) int {
	game := newGame(useJokers)
	for _, line := range data {
		parts := strings.Split(line, " ")
		hand := parts[0]
		bidString := parts[1]
		bid, _ := strconv.Atoi(bidString)
		game.AddHand(hand, bid)
	}
	return game.GetScore()
}

func newGame(useJokers bool) Game {
	g := Game{useJokers: useJokers}
	if useJokers {
		g.rankingSystem = rankingSystemWithJokers
	} else {
		g.rankingSystem = rankingSystem
	}
	return g
}

type Game struct {
	useJokers     bool
	rankingSystem map[string]int
	Hands         []Hand
}

type Hand struct {
	Cards string
	Bid   int
	Rank  int
}

func (g *Game) AddHand(hand string, bid int) {
	h := Hand{Cards: hand, Bid: bid}
	g.Hands = append(g.Hands, h)
}

func (g *Game) RankHands() {
	for i, hand := range g.Hands {
		hand.rank()
		g.Hands[i] = hand
	}
}
func (g *Game) RankHandsWithJokers() {
	for i, hand := range g.Hands {
		hand.rankWithJokers()
		g.Hands[i] = hand
	}
}

func (g *Game) GetScore() int {
	if g.useJokers {
		g.RankHandsWithJokers()
	} else {
		g.RankHands()
	}
	sort.Sort(g)
	multiplier := 1
	total := 0
	for _, hand := range g.Hands {
		score := hand.Bid
		total += score * multiplier
		multiplier++
	}
	return total
}

func (h *Hand) rankWithJokers() {
	possibleHands := []string{h.Cards}
	cardList := "AKQT98765432"
	cardListItems := strings.Split(cardList, "")
	for _, item := range cardListItems {
		new := strings.ReplaceAll(h.Cards, "J", item)
		possibleHands = append(possibleHands, new)
	}
	possibleRanks := []int{}
	for _, hand := range possibleHands {
		rank := getRank(hand)
		possibleRanks = append(possibleRanks, rank)
	}
	sort.Ints(possibleRanks)
	highestRank := possibleRanks[0]
	h.Rank = highestRank
}

func (h *Hand) rank() {
	h.Rank = getRank(h.Cards)
}

func getRank(cards string) int {
	ranker := make(map[rune]int)
	for _, letter := range cards {
		ranker[letter] += 1
	}
	if len(ranker) == 1 {
		return 1
	}
	hasFour := false
	hasThree := false
	hasTwo := false
	twoCount := 0
	for _, item := range ranker {
		if item == 4 {
			hasFour = true
		}
		if item == 3 {
			hasThree = true
		}
		if item == 2 {
			hasTwo = true
			twoCount++
		}
	}
	if hasFour {
		return 2
	}
	if hasThree {
		if hasTwo {
			return 3
		}
		return 4
	}
	if twoCount == 2 { // 2 pair
		return 5
	}
	if hasTwo {
		return 6
	}
	return 7

}

// Implement the Len method required by sort.Interface
func (g Game) Len() int {
	return len(g.Hands)
}

// Implement the Less method required by sort.Interface
func (g Game) Less(i, j int) bool {
	if g.Hands[i].Rank == g.Hands[j].Rank { // If the cards are the same, go card by card
		iStrings := strings.Split(g.Hands[i].Cards, "") // avoid runes :)
		jStrings := strings.Split(g.Hands[j].Cards, "")
		for x := 0; x < len(jStrings); x++ {
			if g.rankingSystem[iStrings[x]] == g.rankingSystem[jStrings[x]] {
				continue
			}
			return g.rankingSystem[iStrings[x]] > g.rankingSystem[jStrings[x]] // smaller wins
		}
	}
	return g.Hands[i].Rank > g.Hands[j].Rank
}

// Implement the Swap method required by sort.Interface
func (g Game) Swap(i, j int) {
	g.Hands[i], g.Hands[j] = g.Hands[j], g.Hands[i]
}

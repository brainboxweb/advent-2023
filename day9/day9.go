package day9

import (
	"strconv"
	"strings"
)

func Part2(data []string) int {
	ret := 0
	i := 0
	for _, line := range data {
		i++
		numbs := []int{}
		items := strings.Split(line, " ")
		for _, item := range items {
			numb, _ := strconv.Atoi(item)
			numbs = append(numbs, numb)
		}
		collection := arrange(numbs)
		extra := extrapolateOne(collection)
		ret += extra
	}

	return ret
}

func Part11(data []string) int {
	ret := 0
	i := 0
	for _, line := range data {
		i++
		numbs := []int{}
		items := strings.Split(line, " ")
		for _, item := range items {
			numb, _ := strconv.Atoi(item)
			numbs = append(numbs, numb)
		}
		collection := arrange(numbs)
		extra := extrapolateTwo(collection)
		ret += extra
	}

	return ret
}

func arrange(numbs []int) [][]int {
	col := [][]int{}
	col = append(col, numbs)
	next := numbs
	for {
		next = getNext(next)
		if next == nil {
			break
		}
		col = append(col, next)
	}

	return col
}

func extrapolateOne(collection [][]int) int {
	ret := 0
	for j := len(collection) - 1; j > -1; j-- {
		first := collection[j][0]
		ret = first - ret
	}

	return ret
}

func extrapolateTwo(collection [][]int) int {
	ret := 0
	for j := len(collection) - 1; j > -1; j-- {
		last := collection[j][len(collection[j])-1]
		ret += last
	}

	return ret
}

func getNext(data []int) []int {
	diffs := []int{}
	for i := range len(data) - 1 {
		diff := data[i+1] - data[i]
		diffs = append(diffs, diff)
	}
	if isAllZeros(diffs) {
		return nil
	}

	return diffs
}

func isAllZeros(data []int) bool {
	for _, item := range data {
		if item != 0 {
			return false
		}
	}

	return true
}

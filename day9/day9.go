package day9

import (
	"strconv"
	"strings"
)

func Day9(data []string, right bool) int {
	i := 0
	ret := 0
	for _, line := range data {
		i++
		numbs := []int{}
		items := strings.Split(line, " ")
		for _, item := range items {
			numb, _ := strconv.Atoi(item)
			numbs = append(numbs, numb)
		}
		collection := arrange(numbs)
		extra := extrapolate(collection, right)
		ret = ret + extra
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

func extrapolate(collection [][]int, right bool) int {
	ret := 0
	if right {
		// any number from last line is the first number
		// next is above + last on ssecond last line
		// etc.
		for j := len(collection) - 1; j > -1; j-- {
			last := collection[j][len(collection[j])-1]
			ret += last
		}
	} else {
		for j := len(collection) - 1; j > -1; j-- {
			first := collection[j][0]
			ret = first - ret /// not sure why!
		}
	}
	return ret
}

func getNext(data []int) []int {
	diffs := []int{}
	for i := 0; i < len(data)-1; i++ {
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

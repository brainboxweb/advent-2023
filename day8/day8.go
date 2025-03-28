package day8

import (
	"strings"
)

type Elements struct {
	Elements map[string]Element
}

type Element struct {
	Left  string
	Right string
}

func newElement(left, right string) Element {
	return Element{Left: left, Right: right}
}

//revive:disable:cognitive-complexity

func Part1(data []string) int {
	instructions, elems := process(data)
	stepCount := 0
	index := "AAA"
	for {
		for _, move := range instructions {
			stepCount++
			if move == "L" {
				index = elems[index].Left
			} else {
				index = elems[index].Right
			}
			if index == "ZZZ" {
				return stepCount
			}
		}
	}
}

//revive:enable:cognitive-complexity

// There's a looping thing going on here.
// - First loops in 2
// - Second loops in 3
// - Therefore they come together in 6: common multiple?
func Part2(data []string) int {
	instructions, elems := process(data)
	targetIndexes := getIndexesBySuffix(elems, "A")
	myChan := make(chan int)
	for _, targetIndex := range targetIndexes {
		go func(targetIndex string) {
			runCommands(elems, instructions, targetIndex, myChan)
		}(targetIndex)
	}
	counts := []int{}
	for i := 0; i < len(targetIndexes); i++ {
		counts = append(counts, <-myChan)
	}
	lcm := LCM(counts[0], counts[1], counts[2:]...)
	return lcm
}

func process(data []string) ([]string, map[string]Element) {
	first := data[0]
	instructions := strings.Split(first, "")
	elems := make(map[string]Element)
	other := data[1:]
	for _, line := range other {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		left := strings.Trim(parts[2], " =(),")
		right := strings.Trim(parts[3], " =(),")
		elem := newElement(left, right)
		elems[strings.Trim(parts[0], " =(),")] = elem
	}
	return instructions, elems
}

func runCommands(
	elems map[string]Element,
	instructions []string,
	targetIndex string,
	myChan chan int,
) {
	stepCount := 0

	for {
		for _, move := range instructions { // One at a time!
			stepCount++
			nextIndex := getNextIndex(elems, targetIndex, move)
			parts := strings.Split(nextIndex, "")
			if parts[2] == "Z" {
				myChan <- stepCount
				return // important
			}
			targetIndex = nextIndex
		}
	}
}

func getNextIndex(elems map[string]Element, index, move string) string {
	if move == "L" {
		return elems[index].Left
	}
	return elems[index].Right
}

func getIndexesBySuffix(elems map[string]Element, suffix string) []string {
	ret := []string{}
	for i := range elems {
		parts := strings.Split(i, "")
		if parts[2] == suffix {
			ret = append(ret, i)
		}
	}
	return ret
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

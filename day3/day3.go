package day3

import (
	"fmt"
)

const charSlash = 47
const charColon = 58
const charPeriod = 46
const charAsterix = 42

//revive:disable:cognitive-complexity
//revive:disable:cyclomatic

func Part1(data []string) int {
	theArray := [][]rune{}
	for _, line := range data {
		lr := []rune{}
		for _, r := range line {
			lr = append(lr, r)
		}
		theArray = append(theArray, lr)
	}
	sum := 0
	for x, line := range theArray {
		inNumber := false
		numberParts := []int{}
		for y, chr := range line {
			if chr > charSlash && chr < charColon {
				numberParts = append(numberParts, int(chr)-48) // that's okay
				inNumber = true
				// Fix the end of the line
				if y == len(line)-1 {
					if isPartNumber(theArray, x, y, len(numberParts)) {
						number := buildNumber(numberParts)
						sum += number
					}
				}
				continue
			}
			// NOT A DIGIT, not at the end of line
			if inNumber {
				if isPartNumber(theArray, x, y, len(numberParts)) {
					number := buildNumber(numberParts)
					sum += number
				}
			}
			// reset
			numberParts = nil
			inNumber = false
		}
	}

	return sum
}

func Part2(data []string) int {
	// an array of the whole thing
	theArray := [][]rune{}
	for _, line := range data {
		lr := []rune{}
		for _, r := range line {
			lr = append(lr, r)
		}
		theArray = append(theArray, lr)
	}

	theStars := make(map[string][]int) // store with a key of the ASTERIX "i_j"

	// sum := 0
	for x, line := range theArray {
		inNumber := false
		numberParts := []int{}
		for y, chr := range line {
			if chr > charSlash && chr < charColon {
				numberParts = append(numberParts, int(chr)-48) // that's okay
				inNumber = true
				// Fix the end of the line
				if y == len(line)-1 {
					stars := getAsterixes(theArray, x, y, len(numberParts))
					if len(stars) > 0 {
						number := buildNumber(numberParts)
						for _, star := range stars {
							index := fmt.Sprint(star[0], " ", star[1])
							theStars[index] = append(theStars[index], number)
						}
					}
				}
				continue
			}

			// NOT A DIGIT, not at the end of line
			if inNumber {
				stars := getAsterixes(theArray, x, y, len(numberParts))
				if len(stars) > 0 {
					number := buildNumber(numberParts)
					for _, star := range stars {
						index := fmt.Sprint(star[0], " ", star[1])
						theStars[index] = append(theStars[index], number)
					}
				}
			}
			// reset
			numberParts = nil
			inNumber = false
		}
	}
	ret := 0
	for _, thing := range theStars {
		if len(thing) == 2 {
			mult := thing[0] * thing[1]
			ret += mult
		}
	}
	return ret
}

func isPartNumber(shit [][]rune, x, y, size int) bool {
	maxI := len(shit) - 1
	maxJ := len(shit[0]) - 1
	for i := (x - 1); i < (x + 2); i++ {
		for j := y - size - 1; j < y+1; j++ {
			if i < 0 || j < 0 {
				continue
			}
			if i > maxI {
				continue
			}
			if j > maxJ {
				continue
			}
			if isSpecialCharacter(shit[i][j]) {
				return true
			}
		}
	}
	return false
}

func getAsterixes(theArray [][]rune, x, y, size int) [][]int {
	maxI := len(theArray) - 1
	maxJ := len(theArray[0]) - 1
	points := [][]int{}
	for i := (x - 1); i < (x + 2); i++ { // don't worry about 0,0
		for j := y - size - 1; j < y+1; j++ {
			if i < 0 || j < 0 {
				continue
			}
			if i > maxI {
				continue
			}
			if j > maxJ {
				continue
			}
			if theArray[i][j] == charAsterix {
				// Store and carry on
				point := []int{i, j}
				points = append(points, point)
			}
		}
	}

	return points
}

//revive:enable:cognitive-complexity

func isSpecialCharacter(chr rune) bool {
	if chr > charSlash && chr < charColon {
		return false
	}
	if chr == charPeriod {
		return false
	}

	return true
}

func buildNumber(numberParts []int) int {
	ret := 0
	multiplier := 1
	for i := len(numberParts) - 1; i >= 0; i-- {
		ret += numberParts[i] * multiplier
		multiplier *= 10
	}
	return ret
}

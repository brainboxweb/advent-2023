package day1

import (
	"regexp"
	"strconv"
)

const colonIndex = 58

func Part1(data []string) int {
	ret := 0
	for _, item := range data {
		ret += DoCalcPart1(item)
	}

	return ret
}

func Part2(data []string) int {
	ret := 0
	for _, item := range data {
		ret += DoCalcPart2(item)
	}

	return ret
}

func DoCalcPart1(in string) int {
	var first, last rune
	haveFirst := false
	for _, char := range in {
		if char < colonIndex {
			if !haveFirst {
				first = char
				haveFirst = true
			}
			last = char // gets overitten
		}
	}
	twofer := string(first) + string(last)
	ret, err := strconv.Atoi(twofer)
	if err != nil {
		panic(err) // not expected
	}

	return ret
}

func DoCalcPart2(in string) int {
	rr := map[int]*regexp.Regexp{
		1: regexp.MustCompile("^(one).*"),
		2: regexp.MustCompile("^(two).*"),
		3: regexp.MustCompile("^(three).*"),
		4: regexp.MustCompile("^(four).*"),
		5: regexp.MustCompile("^(five).*"),
		6: regexp.MustCompile("^(six).*"),
		7: regexp.MustCompile("^(seven).*"),
		8: regexp.MustCompile("^(eight).*"),
		9: regexp.MustCompile("^(nine).*"),
	}

	numbs := []int{}
	for i, char := range in {
		if char < colonIndex {
			numbs = append(numbs, int(char)-48)
			continue
		}
		for number, rx := range rr {
			if rx.FindString(in[i:]) != "" {
				numbs = append(numbs, number)
			}
		}
	}

	return numbs[0]*10 + numbs[len(numbs)-1]
}

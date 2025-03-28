package day1

import (
	"regexp"
	"strconv"
)

func Day1(data []string) int {
	ret := 0
	for _, item := range data {
		ret += DoCalcDay1(item)
	}

	return ret
}

func Day1a(data []string) int {
	ret := 0
	for _, item := range data {
		ret += DoCalcDay1a(item)
	}

	return ret
}

func DoCalcDay1a(in string) int {
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

	numbs := []int{} // for collecting

	for i, char := range in {
		if char < 58 {
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

func DoCalcDay1(in string) int {
	var first, last rune
	haveFirst := false
	for _, char := range in {
		if char < 66 { // is that correct???
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

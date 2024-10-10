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

// func toChar(i int) rune {
// 	return rune('A' - 1 + i)
// }

// func Day1_2(data []int) int {
// 	elves := getElves(data)
// 	sort.Sort(ByFood(elves))
// 	top3 := 0
// 	for i := 0; i < 3; i++ {
// 		top3 += elves[i].total
// 	}

// 	return top3
// }

// func (data []int) []elf {
// 	elves := []elf{}
// 	index := 1
// 	e := elf{id: index}
// 	for _, cals := range data {
// 		if cals == 0 { // write the elf
// 			elves = append(elves, e)
// 			index++
// 			e = elf{id: index}
// 			continue
// 		}
// 		e.addFood(cals)
// 	}
// 	elves = append(elves, e) // write the final elf

// 	return elves
// }

// type elf struct {
// 	id    int
// 	food  []int
// 	total int
// }

// func (e *elf) addFood(f int) {
// 	e.food = append(e.food, f)
// 	e.total += f
// }

// type ByFood []elf

// func (f ByFood) Len() int           { return len(f) }
// func (f ByFood) Less(i, j int) bool { return f[i].total > f[j].total }
// func (f ByFood) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

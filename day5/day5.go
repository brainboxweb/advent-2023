package day5

import (
	"sort"
	"strconv"
	"strings"
	"sync"
)

func Part1(input []string) int {
	lineIndex := []int{}
	for i, line := range input {
		if line == "" {
			lineIndex = append(lineIndex, i)
		}
	}
	seedsString := input[0]
	seedRanges := parseSeedsInput(seedsString)

	return mapping(input, lineIndex, seedRanges)
}

func Part2(input []string) int {
	lineIndex := []int{}
	for i, line := range input {
		if line == "" {
			lineIndex = append(lineIndex, i)
		}
	}
	seedsString := input[0]
	seedRanges := parseSeedsInputPart2(seedsString)

	return mapping(input, lineIndex, seedRanges)
}

func mapping(input []string, lineIndex []int, seedRanges []SeedRange) int {
	seedToSoil := input[lineIndex[0]+2 : lineIndex[1]]
	soilToFertilizer := input[lineIndex[1]+2 : lineIndex[2]]
	fertilizerToWater := input[lineIndex[2]+2 : lineIndex[3]]
	waterToLight := input[lineIndex[3]+2 : lineIndex[4]]
	lightToTemperature := input[lineIndex[4]+2 : lineIndex[5]]
	temperatureToHumidity := input[lineIndex[5]+2 : lineIndex[6]]
	humidityToLocation := input[lineIndex[6]+2:]

	seedToSoilMap := makeMapper(seedToSoil)
	soilToFertilizerMap := makeMapper(soilToFertilizer)
	fertilizerToWaterMap := makeMapper(fertilizerToWater)
	waterToLightMap := makeMapper(waterToLight)
	lightToTemperatureMap := makeMapper(lightToTemperature)
	temperatureToHumidityMap := makeMapper(temperatureToHumidity)
	humidityToLocationMap := makeMapper(humidityToLocation)

	worker := func(rang SeedRange, c chan int) {
		lowest := 1000000000000000000
		start := rang.Min
		for i := 0; i < rang.Count+1; i++ { // stuill not sure!!
			soil := seedToSoilMap.Translate(start + i)
			fert := soilToFertilizerMap.Translate(soil)
			water := fertilizerToWaterMap.Translate(fert)
			light := waterToLightMap.Translate(water)
			temp := lightToTemperatureMap.Translate(light)
			hum := temperatureToHumidityMap.Translate(temp)
			locn := humidityToLocationMap.Translate(hum)
			if locn < lowest {
				lowest = locn
			}
		}
		c <- lowest
	}

	c := make(chan int)

	var wg sync.WaitGroup
	for _, rang := range seedRanges {
		wg.Add(1)
		go func(rang SeedRange) {
			defer wg.Done()
			worker(rang, c)
		}(rang)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	locations := []int{}
	for loc := range c {
		locations = append(locations, loc)
	}

	sort.Ints(locations)
	return locations[0]
}

func parseSeedsInput(input string) []SeedRange { // to align with part 2
	elements := strings.Split(input, " ")
	// ditch the first element
	elements = elements[1:]
	ret := []SeedRange{}
	for _, el := range elements {
		numb, _ := strconv.Atoi(el)
		rang := SeedRange{numb, 1} // Alway 1 in this case
		ret = append(ret, rang)
	}
	return ret
}

type SeedRange struct {
	Min   int
	Count int
}

func parseSeedsInputPart2(input string) []SeedRange {
	elements := strings.Split(input, " ")
	// ditch the first element
	elements = elements[1:]
	items := []int{}
	for _, el := range elements {
		numb, _ := strconv.Atoi(el)
		items = append(items, numb)
	}
	ranges := []SeedRange{}
	for i := 0; i < len(items)-1; i += 2 {
		minimum := items[i]
		rang := SeedRange{Min: minimum, Count: items[i+1]}
		ranges = append(ranges, rang)
	}

	return ranges
}

func makeMapper(input []string) Mapper {
	mapper := Mapper{}
	for _, line := range input {
		elements := strings.Split(line, " ")
		source, _ := strconv.Atoi(elements[1])
		dest, _ := strconv.Atoi(elements[0])
		rang, _ := strconv.Atoi(elements[2])
		part := MapperPart{Source: source, Dest: dest, Rang: rang}
		mapper.Parts = append(mapper.Parts, part)
	}

	return mapper
}

type Mapper struct {
	Parts []MapperPart
}

type MapperPart struct {
	Source int
	Dest   int
	Rang   int
}

func (m *Mapper) Translate(input int) int {
	for _, part := range m.Parts {
		minimum := part.Source
		maximum := minimum + part.Rang - 1 // If length is one, min==max
		if input >= minimum && input <= maximum {
			return part.Dest + input - minimum // mapping found
		}
	}

	return input // Default: just return the input.
}

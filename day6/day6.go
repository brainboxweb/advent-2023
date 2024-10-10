package day6

func Day6(races []Race) int {

	winCounts := []int{}
	for _, race := range races {
		winCounts = append(winCounts, race.DoIt())
	}
	ret := 1
	for _, thing := range winCounts {
		ret *= thing
	}
	return ret
}

type Race struct {
	Time   int
	Record int
}

func (r Race) DoIt() int {
	winCount := 0
	for i := 1; i < r.Time; i++ {
		distance := (r.Time - i) * i
		if distance > r.Record {
			winCount++
		}
	}

	return winCount
}

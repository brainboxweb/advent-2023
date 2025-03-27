package game

func New(index int) *Game {
	return &Game{index: index}
}

type Game struct {
	index int
	turns []Turn
}

func (g *Game) Index() int {
	return g.index
}

func (g *Game) AddTurn(turnData map[string]int) { // color, count
	t := NewTurn()
	for color, count := range turnData {
		t.addCubes(color, count)
	}
	g.turns = append(g.turns, t)
}

func (g *Game) IsPossible(contents map[string]int) bool {
	colors := []string{"red", "green", "blue"}
	for _, color := range colors {
		total := contents[color]
		actuals := g.GetColorCounts(color)
		for _, actual := range actuals {
			if actual > total {
				return false
			}
		}
	}

	return true
}

func (g *Game) GetColorCounts(color string) []int { // to unexp
	counts := []int{}
	for _, t := range g.turns {
		counts = append(counts, t.cubes[color])
	}

	return counts
}

type Turn struct {
	cubes map[string]int
}

func NewTurn() Turn {
	t := Turn{}
	t.cubes = make(map[string]int)
	return t
}

func (t *Turn) addCubes(color string, count int) {
	t.cubes[color] += count
}

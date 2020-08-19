package kata

type game struct {
	score int
}

func (g *game) Roll(amount int) {
	g.score += amount
}

func (g *game) Score() int {
	return g.score
}

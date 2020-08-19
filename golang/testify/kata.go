package kata

type Game struct {
	score int
}

func (g *Game) Roll(pins int) {
	g.score += pins
}

func (g *Game) Score() int {
	return g.score
}

package kata

type game struct {
	score int
}

func (g *game) Roll(amount int) {
	if amount > 10 || amount < 0 {
		return
	}

	g.score += amount
}

func (g *game) Score() int {
	return g.score
}

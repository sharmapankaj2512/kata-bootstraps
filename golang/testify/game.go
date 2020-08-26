package kata

type game struct {
	score int
	strike bool
}

func (g *game) Roll(amount int) {
	if amount > 10 || amount < 0 {
		return
	}

	if amount == 10 {
		g.strike = true
		return
	}

	g.score += amount
}

func (g *game) Score() int {
	return g.score
}

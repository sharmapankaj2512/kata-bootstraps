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

	if g.strike {
		amount *= 2
		amount += 10
	}
	
	g.score += amount
}

func (g *game) Score() int {
	return g.score
}

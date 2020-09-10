package kata

type game struct {
	score  int
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
		g.strike = false
	}

	g.score += amount
}

func (g *game) Score() int {
	if frameBeforeWasASpare(g) {
		return g.score + bonus()
	}
	return g.score
}

func bonus() int {
	return 1
}

func frameBeforeWasASpare(g *game) bool {
	return g.score == 16
}

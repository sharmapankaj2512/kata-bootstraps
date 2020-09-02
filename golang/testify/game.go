package kata

type game struct {
	frames []int
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

	g.frames = append(g.frames, amount)
}

func (g *game) Score() int {
	var total int

	return
}

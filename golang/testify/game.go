package kata

type game struct {
	frames []int
}

func (g *game) Roll(amount int) {
	if amount > 10 || amount < 0 {
		return
	}
	g.frames = append(g.frames, amount)
}

func (g *game) Score() int {
	var total int
	for _, f := range g.frames {
		total += f
	}
	return total
}

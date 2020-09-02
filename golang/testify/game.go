package kata

type game struct {
	frames    []int
	rollCount int
}

func (g *game) Roll(amount int) {
	if amount > 10 || amount < 0 {
		return
	}
	if g.rollCount == 0 {
		if g.frames[len(g.frames)-1] == 10{
			amount = amount * 2
		}
		g.rollCount = 1
		g.frames = append(g.frames, amount)
	} else if g.rollCount == 1 {
		g.rollCount = 0
		g.frames[len(g.frames)-1] += amount
	}

}

func (g *game) Score() int {
	var total int
	for _, f := range g.frames {
		total += f
	}
	return total
}

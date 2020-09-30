package kata

type game struct{
	field [][]byte
	iterations int
}

func (g *game) iterate() {
	g.iterations++
}



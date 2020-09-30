package kata

type game struct{
	field [][]byte
	iterations int
}

func (g *game) get(x, y int) byte {
	if x < 0 || x >= 3 {
		return 0
	}

	if y < 0 || y >= 3 {
		return 0
	}

	return g.field[x][y]
}

func (g *game) iterate() {
	g.iterations++

	for i, row := range g.field {
		for j := range row {
			cells := g.get(i-1, j-1)
			cells += g.get(i+0, j-1)
			cells += g.get(i+1, j-1)
			cells += g.get(i-1, j-0)
			cells += g.get(i+1, j-0)
			cells += g.get(i-1, j+1)
			cells += g.get(i+0, j+1)
			cells += g.get(i+1, j+1)

			switch cells {
			case 0:
				g.field[i][j] = 0
			}
		}
	}
}



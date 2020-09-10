package kata

type game struct {
	score  int
	strike bool
	rolls  []int
}

func (game *game) Roll(fallenPins int) {
	game.rolls = append(game.rolls, fallenPins)

	if fallenPins > 10 || fallenPins < 0 {
		return
	}

	if fallenPins == 10 {
		game.strike = true
		return
	}

	if game.strike {
		fallenPins *= 2
		game.strike = false
	}

	game.score += fallenPins
}

func (game *game) Score() int {
	if frameBeforeWasASpare(game) {
		return game.score + bonus()
	}
	return game.score
}

func bonus() int {
	return 1
}

func frameBeforeWasASpare(g *game) bool {
	return g.score == 16
}

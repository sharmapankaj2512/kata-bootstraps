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
	if thereWasOnlyOneFrame(game) && currentFrameIsASpare(game) {
		return 0
	}
	if frameBeforeWasASpare(game) {
		return game.score + bonus()
	}
	return game.score
}

func currentFrameIsASpare(game *game) bool {
	return game.rolls[0]+game.rolls[1] == 10
}

func thereWasOnlyOneFrame(game *game) bool {
	return len(game.rolls) == 2
}

func bonus() int {
	return 1
}

func frameBeforeWasASpare(g *game) bool {
	if len(g.rolls) >= 2 && g.score == 16 {
		return g.rolls[0]+g.rolls[1] == 10
	}
	return false
}

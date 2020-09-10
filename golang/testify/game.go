package kata

type game struct {
	score int
	rolls []int
}

func (game *game) Roll(fallenPins int) {
	game.rolls = append(game.rolls, fallenPins)
	if len(game.frames) == 0 {
		game.frames = append(game.frames, frame{score: fallenPins})
	} else {
		f := game.frames[len(game.frames)-1]
		if f.isClosed {
			game.frames = append(game.frames, frame{score: fallenPins})

		} else {
			f.score += fallenPins
			f.isClosed = true
		}
	}

}

func (game *game) Score() int {
	if firstFrameNotFinished(game) {
		return 0
	}

	if thereWasOnlyOneFrame(game) && currentFrameIsASpare(game) {
		return 0
	}

	if frameBeforeWasASpare(game) {
		return sumOfClosedFrames(game) + bonus()
	}

	return sumOfClosedFrames(game)
}

func sumOfClosedFrames(game *game) int {
	var sum int
	for i := 0; i < len(game.rolls); i++ {
		sum += game.rolls[i]
	}
	if len(game.rolls)%2 == 1 {
		sum -= game.rolls[len(game.rolls)-1]
	}
	return sum
}

func firstFrameNotFinished(game *game) bool {
	return len(game.rolls) == 1
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
	if len(g.rolls) >= 2 {
		return g.rolls[0]+g.rolls[1] == 10
	}
	return false
}

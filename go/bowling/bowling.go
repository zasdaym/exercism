package bowling

import (
	"errors"
)

const (
	maxPins       = 10
	framesPerGame = 10
	pinsPerFrame  = 10
	maxFrameRolls = 2
	maxFinalRolls = 3
)

var (
	ErrInvalidPinCount     = errors.New("invalid pin count")
	ErrGameAlreadyFinished = errors.New("game is already finished")
	ErrGameNotFinished     = errors.New("game is not finished")
)

type Frame []int

func (f Frame) total() int {
	var sum int
	for _, pins := range f {
		sum += pins
	}
	return sum
}

func (f Frame) isStrike() bool {
	return len(f) > 0 && f[0] == maxPins
}

func (f Frame) isSpare() bool {
	return len(f) > 1 && f[0]+f[1] == pinsPerFrame
}

type Game struct {
	currentFrame Frame   // Current frame being played
	frames       []Frame // Completed frames
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Roll(pins int) error {
	if g.isFinished() {
		return ErrGameAlreadyFinished
	}
	if pins < 0 || pins > maxPins {
		return ErrInvalidPinCount
	}

	g.currentFrame = append(g.currentFrame, pins)

	if len(g.frames) == framesPerGame-1 {
		return g.handleFinalFrame()
	}

	if g.currentFrame.total() > pinsPerFrame {
		return ErrInvalidPinCount
	}
	if pins == maxPins || len(g.currentFrame) == maxFrameRolls {
		g.submitCurrentFrame()
	}
	return nil
}

func (g *Game) Score() (int, error) {
	if !g.isFinished() {
		return 0, ErrGameNotFinished
	}

	var rolls []int
	for _, frame := range g.frames {
		rolls = append(rolls, frame...)
	}

	var (
		result int
		roll   int
	)

	for frame := 0; frame < framesPerGame; frame++ {
		if roll >= len(rolls) {
			break
		}

		switch {
		case rolls[roll] == maxPins:
			result += maxPins + rolls[roll+1] + rolls[roll+2]
			roll++
		case rolls[roll]+rolls[roll+1] == pinsPerFrame:
			result += pinsPerFrame + rolls[roll+2]
			roll += maxFrameRolls
		default:
			result += rolls[roll] + rolls[roll+1]
			roll += maxFrameRolls
		}
	}

	return result, nil
}

func (g *Game) handleFinalFrame() error {
	switch {
	case g.currentFrame.isStrike():
		if len(g.currentFrame) == maxFinalRolls {
			if g.currentFrame[1] != maxPins && g.currentFrame[1]+g.currentFrame[2] > pinsPerFrame {
				return ErrInvalidPinCount
			}
			g.submitCurrentFrame()
		}
	case g.currentFrame.isSpare():
		if len(g.currentFrame) == maxFinalRolls {
			g.submitCurrentFrame()
		}
	default:
		if len(g.currentFrame) == maxFrameRolls {
			g.submitCurrentFrame()
		}
	}
	return nil
}

func (g *Game) isFinished() bool {
	return len(g.frames) == framesPerGame
}

func (g *Game) submitCurrentFrame() {
	g.frames = append(g.frames, g.currentFrame)
	g.currentFrame = nil
}

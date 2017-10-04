package bowling

import (
	"errors"
	"fmt"
)

const testVersion = 1

type Game struct {
	frames                [10][3]int
	frameIndex, rollIndex int
}

func NewGame() *Game {
	return &Game{}
}

func (g Game) String() string {
	return fmt.Sprintf(
		"{frames:%v frameIndex:%v rollIndex:%v}",
		g.frames, g.frameIndex, g.rollIndex)
}

func (g *Game) Roll(pins int) error {
	if err := applyRules(g.frames[g.frameIndex], g.frameIndex, g.rollIndex, pins); err != nil {
		return err
	}

	g.frames[g.frameIndex][g.rollIndex] = pins
	g.rollIndex++

	if g.frameIndex < 9 && g.rollIndex > 1 {
		g.rollIndex = 0
		g.frameIndex++
	}

	return nil
}

func applyRules(frame [3]int, frameIndex, rollIndex, pins int) error {
	const (
		minPins         = 0
		maxPins         = 10
		finalFrameIndex = 9
	)
	switch {
	case pins < minPins:
		return errors.New("Can't roll a negative number")
	case pins > maxPins:
		return errors.New("Can't roll more than 10")
	case frameIndex < finalFrameIndex && sum(frame)+pins > 10:
		return errors.New("Can't roll more than 10 in a single frame")
	case frameIndex == finalFrameIndex && rollIndex == 2 && frame[1]+pins > maxPins:
		return errors.New("Can't roll more than 10 in the last two rolls of the 10th frame")
	case frameIndex == finalFrameIndex && rollIndex == 2:
		return errors.New("The game is over. No more rolls are allowed.")
	}
	return nil
}

func (g Game) Score() (score int, e error) {
	for _, frame := range g.frames {
		score += sum(frame)
		// TODO Resume scoring here
	}
	return score, nil
}

func sum(nums [3]int) (total int) {
	for _, n := range nums {
		total += n
	}
	return total
}

package bowling

import (
	"errors"
)

const testVersion = 1

const (
	minPins         = 0
	maxPins         = 10
	finalFrameIndex = 9
)

// Game stores bowling activity and scores the result.
type Game struct {
	frames                [10][3]int
	frameIndex, rollIndex int
}

// NewGame creates a Game for scoring bowling.
func NewGame() *Game {
	return &Game{}
}

// Roll records a throw of the ball.
func (g *Game) Roll(pins int) error {
	if err := g.applyRules(g.frames[g.frameIndex], g.frameIndex, g.rollIndex, pins); err != nil {
		return err
	}

	g.frames[g.frameIndex][g.rollIndex] = pins
	g.rollIndex++

	if g.frameIndex < finalFrameIndex && isStrike(g.frames[g.frameIndex]) {
		g.rollIndex = 0
		g.frameIndex++
	} else if g.frameIndex < finalFrameIndex && g.rollIndex > 1 {
		g.rollIndex = 0
		g.frameIndex++
	}

	return nil
}

func (g Game) applyRules(frame [3]int, frameIndex, rollIndex, pins int) error {
	switch {
	case pins < minPins:
		return errors.New("can't roll a negative number")
	case pins > maxPins:
		return errors.New("can't roll more than 10")
	case frameIndex < finalFrameIndex && sum(frame)+pins > maxPins:
		return errors.New("can't roll more than 10 in a single frame")
	case frameIndex == finalFrameIndex && rollIndex == 2:
		if isSpare(frame) {
			return nil
		} else if isStrike(frame) && frame[1] == maxPins {
			return nil
		} else if frame[1]+pins > maxPins {
			return errors.New(
				"can't roll more than 10 in the last two rolls of the 10th frame")
		} else if frame[0]+frame[1] < maxPins {
			return errors.New("the game is over; no more rolls are allowed")
		}
	}
	return nil
}

// Score calculates the number of points achived for all rolls.
func (g Game) Score() (score int, e error) {
	if !g.isGameDone() {
		return 0, errors.New("game cannot be scored until it is over")
	}
	for i, frame := range g.frames {
		if i == len(g.frames)-1 {
			// Final frame
			if frame == [3]int{maxPins, maxPins, maxPins} {
				if isStrike(g.frames[i-1]) {
					score += 20
				} else {
					score += 30
				}
			} else {
				score += sum(frame)
			}
		} else {
			// All frames but the final
			score += sum(frame)
			if isSpare(frame) {
				score += g.frames[i+1][0]
			} else if isStrike(frame) {
				score += g.sumOfNextTwoRolls(i)
			}
		}
	}
	return score, nil
}

func (g Game) isGameDone() bool {
	if g.frameIndex < finalFrameIndex {
		return false
	}
	finalFrame := g.frames[finalFrameIndex]
	switch {
	case isOpen(finalFrame) && g.rollIndex == 2:
		return true
	case (isStrike(finalFrame) || isSpare(finalFrame)) && g.rollIndex == 3:
		return true
	}
	return false
}

// sum calculates the simple total of a single frame.
func sum(nums [3]int) (total int) {
	for _, n := range nums {
		total += n
	}
	return total
}

func isSpare(nums [3]int) bool {
	return nums[0] < maxPins && nums[0]+nums[1] == maxPins
}

func isOpen(nums [3]int) bool {
	return nums[0] < maxPins && nums[0]+nums[1] < maxPins
}

func isStrike(nums [3]int) bool {
	return nums[0] == maxPins
}

func (g Game) sumOfNextTwoRolls(index int) (total int) {
	i := index + 1
	if isStrike(g.frames[i]) && i < len(g.frames)-1 {
		// Sum of this strike plus the next roll
		return 10 + g.frames[i+1][0]
	}
	return sum(g.frames[i])
}

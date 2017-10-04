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

    if g.frameIndex < 9 && isStrike(g.frames[g.frameIndex]) {
      g.rollIndex = 0
      g.frameIndex++
    } else if g.frameIndex < 9 && g.rollIndex > 1 {
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
    case frameIndex == finalFrameIndex && rollIndex == 2:
      if isSpare(frame) {
        return nil
      } else if isStrike(frame) && frame[1] == 10 {
        return nil
      } else if frame[1]+pins > maxPins {
        return errors.New("Can't roll more than 10 in the last two rolls of the 10th frame")
      } else if frame[0]+frame[1] < 10 {
        return errors.New("The game is over. No more rolls are allowed.")
      }
    }
    return nil
  }

  func (g Game) Score() (score int, e error) {
    for i, frame := range g.frames {
      score += sum(frame)
      if i < len(g.frames)-1 {
        if isSpare(frame) {
          score += g.frames[i+1][0]
        } else if isStrike(frame) {
          score += g.sumOfNextTwoRolls(i)
        }
      }
    }
    return score, nil
  }

  func sum(nums [3]int) (total int) {
    for _, n := range nums {
      total += n
    }
    return total
  }

  func isSpare(nums [3]int) bool {
    if nums[0] < 10 && nums[0]+nums[1] == 10 {
      return true
    }
    return false
  }

  func isOpen(nums [3]int) bool {
    if nums[0] < 10 && nums[0]+nums[1] < 10 {
      return true
    }
    return false
  }

  func isStrike(nums [3]int) bool {
    if nums[0] == 10 {
      return true
    }
    return false
  }

  func (g Game) sumOfNextTwoRolls(index int) (total int) {
    for i := index+1; i < len(g.frames)-1; i++ {
      if isStrike(g.frames[i]) {
        // Sum of this strike plus the next roll
        return 10 + g.frames[i+1][0]
      } else {
        return sum(g.frames[i])
      }
    }
    return total
  }

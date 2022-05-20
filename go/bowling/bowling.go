package bowling

import (
	"fmt"
)

type Game struct {
	frames []*Frame
	rolls  []int
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Roll(pins int) error {
	if pins < 0 {
		return fmt.Errorf("negative roll is invalid")
	}
	if pins > 10 {
		return fmt.Errorf("roll can not hit more than 10 pins")
	}
	g.rolls = append(g.rolls, pins)
	return g.parseFrames()
}

func (g *Game) parseFrames() error {
	g.frames = []*Frame{}
	for _, roll := range g.rolls {
		isCompleteFrame := roll == 10

		if len(g.frames) >= 10 && g.frames[9].isComplete && !g.frames[9].isStrike() && !g.frames[9].isSpare() {
			return fmt.Errorf("cannot roll after game is over")
		} else if len(g.frames) > 10 && g.frames[9].isComplete && g.frames[9].isSpare() {
			return fmt.Errorf("cannot roll after bonus roll for spare")
		} else if len(g.frames) == 11 && g.frames[9].isComplete && g.frames[9].isStrike() && g.frames[10].isComplete && !g.frames[10].isStrike() {
			return fmt.Errorf("cannot roll after bonus roll for strike")
		} else if len(g.frames) == 0 {
			g.frames = append(g.frames, &Frame{roll, 0, isCompleteFrame})
		} else if !g.lastFrame().isComplete {
			if g.lastFrame().rollOne+roll > 10 {
				return fmt.Errorf("second roll hit more than number of pins on the lane")
			}
			g.lastFrame().rollTwo = roll
			g.lastFrame().isComplete = true
		} else {
			g.frames = append(g.frames, &Frame{roll, 0, isCompleteFrame})
		}
	}
	return nil
}

func (g *Game) Score() (total int, err error) {
	g.parseFrames()

	if len(g.frames) < 10 {
		return 0, fmt.Errorf("not enough frames %v", g.frames)
	}

	for i, frame := range g.frames[0:10] {
		score, err := frame.score(g.frames[i+1:])
		if err != nil {
			return 0, err
		}
		total += score
	}
	return total, nil
}

func (g *Game) lastFrame() *Frame {
	return g.frames[len(g.frames)-1]
}

type Frame struct {
	rollOne    int
	rollTwo    int
	isComplete bool
}

func (f *Frame) String() string {
	return fmt.Sprintf("[%d, %d]", f.rollOne, f.rollTwo)
}

func (f *Frame) score(nextFrames []*Frame) (score int, err error) {
	if f.isOpenFrame() {
		return f.rollOne + f.rollTwo, nil
	}
	nextRolls := nextRolls(nextFrames)
	if f.isSpare() {
		if len(nextRolls) < 1 {
			return 0, fmt.Errorf("not enough rolls after spare")
		}
		return 10 + nextRolls[0], nil
	}
	if f.isStrike() {
		if len(nextRolls) < 2 {
			return 0, fmt.Errorf("not enough rolls after strike")
		}
		return 10 + nextRolls[0] + nextRolls[1], nil
	}
	return 0, fmt.Errorf("frame %v is not an open frame, spare, or strike", f)
}

func (f *Frame) isStrike() bool {
	return isStrike(f.rollOne)
}

func (f *Frame) isSpare() bool {
	return !f.isStrike() && f.rollOne+f.rollTwo == 10
}

func (f *Frame) isOpenFrame() bool {
	return !f.isStrike() && !f.isSpare()
}

func isStrike(roll int) bool {
	return roll == 10
}

func nextRolls(nextFrames []*Frame) (rolls []int) {
	rolls = []int{}
	for _, frame := range nextFrames {
		if frame.isStrike() {
			rolls = append(rolls, frame.rollOne)
		} else {
			rolls = append(rolls, frame.rollOne)
			rolls = append(rolls, frame.rollTwo)
		}
	}
	return rolls
}

package bowling

import (
	"fmt"
	"log"
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
	return nil
}

func (g *Game) parseFrames() {
	g.frames = []*Frame{}
	for _, roll := range g.rolls {
		if len(g.frames) == 0 {
			isComplete := roll == 10
			g.frames = append(g.frames, &Frame{roll, 0, isComplete})
		} else if !g.lastFrame().isComplete {
			g.lastFrame().rollTwo = roll
			g.lastFrame().isComplete = true
		} else {
			isComplete := roll == 10
			g.frames = append(g.frames, &Frame{roll, 0, isComplete})
		}
	}
}

func (g *Game) Score() (total int, err error) {
	g.parseFrames()
	fmt.Printf("rolls: %v\nframes: %v\n", g.rolls, g.frames)
	if len(g.frames) < 10 {
		return 0, fmt.Errorf("not enough frames %v", g.frames)
	}
	for i, frame := range g.frames[0:10] {
		score := frame.score(g.frames[i+1:])
		fmt.Printf("frame: %v score %v\n", frame, score)
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

func (f *Frame) score(nextFrames []*Frame) int {
	if f.isOpenFrame() {
		return f.rollOne + f.rollTwo
	}
	nextRoll, nextNextRoll := nextRolls(nextFrames)
	if f.isSpare() {
		return 10 + nextRoll
	}
	if f.isStrike() {
		return 10 + nextRoll + nextNextRoll
	}
	log.Fatalf("frame %v is not an open frame, spare, or strike", f)
	return 0
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

func nextRolls(nextFrames []*Frame) (nextRoll int, nextNextRoll int) {
	if len(nextFrames) == 0 {
		return 0, 0
	}
	rolls := []int{}
	for _, frame := range nextFrames {
		if frame.isStrike() {
			rolls = append(rolls, frame.rollOne)
		} else {
			rolls = append(rolls, frame.rollOne)
			rolls = append(rolls, frame.rollTwo)
		}
	}

	if len(rolls) >= 2 {
		return rolls[0], rolls[1]
	} else if len(rolls) == 1 {
		return rolls[0], 0
	} else {
		log.Fatalf("rolls %v is empty", rolls)
		return 0, 0
	}
}

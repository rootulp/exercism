package robot

import (
	"fmt"
	"log"
)

// Step 1
type Dir int

const N Dir = 0
const E Dir = 1
const S Dir = 2
const W Dir = 3

var Step1Robot struct {
	X, Y int
	Dir
}

var _ fmt.Stringer = Dir(1729)

func Right() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = E
	case E:
		Step1Robot.Dir = S
	case S:
		Step1Robot.Dir = W
	case W:
		Step1Robot.Dir = N
	default:
		log.Fatalf("unrecognized direction %d", Step1Robot.Dir)
	}
}

func Left() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = W
	case E:
		Step1Robot.Dir = N
	case S:
		Step1Robot.Dir = E
	case W:
		Step1Robot.Dir = S
	default:
		log.Fatalf("unrecognized direction %d", Step1Robot.Dir)
	}
}

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case E:
		Step1Robot.X++
	case S:
		Step1Robot.Y--
	case W:
		Step1Robot.X--
	default:
		log.Fatalf("unrecognized direction %d", Step1Robot.Dir)
	}
}

func (d Dir) String() string {
	switch d {
	case N:
		return "N"
	case E:
		return "E"
	case S:
		return "S"
	case W:
		return "W"
	default:
		return fmt.Sprintf("unrecognized direction %d", d)
	}
}

// Step 2
type Command byte // valid values are 'R', 'L', 'A'
type RU int
type Pos struct{ X, Y RU }
type Rect struct{ Min, Max Pos }
type Action uint
type Step2Robot struct {
	Dir
	Pos
}

func (s *Step2Robot) Initialize(X, Y RU) {
	log.Printf("Initializing robot at %d,%d facing %s", s.Pos.X, s.Pos.Y, s.Dir)
	s.Pos.X = X
	s.Pos.Y = Y
}

func (s *Step2Robot) RotateRight() {
	switch s.Dir {
	case N:
		s.Dir = E
	case E:
		s.Dir = S
	case S:
		s.Dir = W
	case W:
		s.Dir = N
	default:
		log.Fatalf("unrecognized direction %d", s.Dir)
	}
}

func (s *Step2Robot) RotateLeft() {
	switch s.Dir {
	case N:
		s.Dir = W
	case E:
		s.Dir = N
	case S:
		s.Dir = E
	case W:
		s.Dir = S
	default:
		log.Fatalf("unrecognized direction %d", s.Dir)
	}
}

func StartRobot(commands chan Command, actions chan Action) {
	for cmd := range commands {
		actions <- Action(cmd)
	}
	close(actions)
}

func Room(extent Rect, robot Step2Robot, actions chan Action, report chan Step2Robot) {
	for action := range actions {
		switch action {
		case 'R':
			robot.RotateRight()
		case 'L':
			robot.RotateLeft()
		case 'A':
			if robot.Dir == N && robot.Pos.Y < extent.Max.Y {
				robot.Pos.Y++
			} else if robot.Dir == E && robot.Pos.X < extent.Max.X {
				robot.Pos.X++
			} else if robot.Dir == S && robot.Pos.Y > extent.Min.Y {
				robot.Pos.Y--
			} else if robot.Dir == W && robot.Pos.X > extent.Min.X {
				robot.Pos.X--
			}
		default:
			log.Fatalf("unrecognized action %d", action)
		}
	}
	report <- robot
	close(report)
}

// Step 3
type Step3Robot struct {
	Name string
	Step2Robot
}
type Action3 struct{}

func StartRobot3(name, script string, action chan Action3, log chan string) {
	panic("Please implement the StartRobot3 function")
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	panic("Please implement the Room3 function")
}

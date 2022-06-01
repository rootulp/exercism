package robot

import (
	"fmt"
	"log"
)

// Step 1
const N Dir = 0
const E Dir = 1
const S Dir = 2
const W Dir = 3

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
type Action uint

func (s *Step2Robot) Initialize(X, Y RU) {
	log.Printf("Initializing robot at %d,%d facing %s", s.Pos.Easting, s.Pos.Northing, s.Dir)
	s.Pos.Easting = X
	s.Pos.Northing = Y
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

func StartRobot(commands <-chan Command, actions chan<- Action) {
	for command := range commands {
		actions <- Action(command)
	}
	close(actions)
}

func Room(extent Rect, robot Step2Robot, actions <-chan Action, report chan<- Step2Robot) {
	for action := range actions {
		switch action {
		case 'R':
			robot.RotateRight()
		case 'L':
			robot.RotateLeft()
		case 'A':
			if robot.Dir == N && robot.Pos.Northing < extent.Max.Northing {
				robot.Pos.Northing++
			} else if robot.Dir == E && robot.Pos.Easting < extent.Max.Easting {
				robot.Pos.Easting++
			} else if robot.Dir == S && robot.Pos.Northing > extent.Min.Northing {
				robot.Pos.Northing--
			} else if robot.Dir == W && robot.Pos.Easting > extent.Min.Easting {
				robot.Pos.Easting--
			}
		default:
			log.Fatalf("unrecognized action %d", action)
		}
	}
	report <- robot
	close(report)
}

// Step 3
type Action3 struct {
	name   string
	script string
}

func isCommand(cmd rune) bool {
	return cmd == 'A' || cmd == 'R' || cmd == 'L'
}

func StartRobot3(name, script string, actions chan<- Action3, logs chan<- string) {
	for _, command := range script {
		if !isCommand(command) {
			logs <- "invalid script"
			actions <- Action3{name: name, script: ""}
			return
		}
	}
	actions <- Action3{name, script}
}

func Room3(extent Rect, robots []Step3Robot, actions chan Action3, rep chan []Step3Robot, log chan string) {
	for index, robot := range robots {
		if robot.Name == "" {
			log <- "robot has no name"
		}
		if robot.Northing < extent.Min.Northing || robot.Northing > extent.Max.Northing || robot.Easting < extent.Min.Easting || robot.Easting > extent.Max.Easting {
			log <- "robot placed outside room"
		}
		for i := index + 1; i < len(robots); i++ {
			if robot.Name == robots[i].Name {
				log <- fmt.Sprintf("duplicate name: %s", robot.Name)
			}
			if robot.Pos == robots[i].Pos {
				log <- "robots placed in same position"
			}
		}
	}

	var action Action3
	for range robots {
		action = <-actions
		isRobot := false
		for robotIndex, robot := range robots {
			if robot.Name == action.name {
				isRobot = true
				for _, command := range action.script {
					if command == 'R' {
						robot.Dir = (robot.Dir + 1) % 4
					} else if command == 'L' {
						robot.Dir = (robot.Dir + 3) % 4
					} else if command == 'A' {
						if IsBlockedByRobot(robot.Pos, robot.Dir, robots) {
							log <- "running into robot"
							continue
						}
						var moved bool
						robot, moved = HitWallOrMove(robot, extent)
						if !moved {
							log <- "running into wall"
						}
					}
					robots[robotIndex] = robot
				}
			}
		}
		if !isRobot {
			log <- fmt.Sprintf("no robot with name %s exists", action.name)
		}
	}
	rep <- robots
	close(rep)
}

func IsBlockedByRobot(position Pos, direction Dir, robots []Step3Robot) bool {
	newPos := position
	switch direction {
	case N:
		newPos.Northing++
	case E:
		newPos.Easting++
	case S:
		newPos.Northing--
	case W:
		newPos.Easting--
	}
	for _, robot := range robots {
		if robot.Pos == newPos {
			return true
		}
	}
	return false
}

func HitWallOrMove(robot Step3Robot, extent Rect) (Step3Robot, bool) {
	switch robot.Dir {
	case N:
		if robot.Pos.Northing >= extent.Max.Northing {
			return robot, false
		} else {
			robot.Pos.Northing++
		}
	case E:
		if robot.Pos.Easting >= extent.Max.Easting {
			return robot, false
		} else {
			robot.Pos.Easting++
		}
	case S:
		if robot.Pos.Northing <= extent.Min.Northing {
			return robot, false
		} else {
			robot.Pos.Northing--
		}
	case W:
		if robot.Pos.Easting <= extent.Min.Easting {
			return robot, false
		} else {
			robot.Pos.Easting--
		}
	}
	return robot, true
}

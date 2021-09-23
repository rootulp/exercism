package elon

const FULL_BATTERY = 100

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

// Track implements a race track.
type Track struct {
	distance int
}

// CreateCar creates a new car with given specifications.
func CreateCar(speed, batteryDrain int) *Car {
	return &Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      FULL_BATTERY,
		distance:     0,
	}
}

// CreateTrack creates a new track with given distance.
func CreateTrack(distance int) Track {
	panic("Please implement CreateTrack() method")
}

// Drive drives the car one time.
func (car *Car) Drive() {
	panic("Please implement Drive() method")
}

// CanFinish checks if a car is able to finish a certain track.
func (car *Car) CanFinish(track Track) bool {
	panic("Please implement CanFinish() method")
}

// DisplayDistance displays the distance the car is driven.
func (car *Car) DisplayDistance() string {
	panic("Please implement DisplayDistance() method")
}

// DisplayBattery displays the battery level.
func (car *Car) DisplayBattery() string {
	panic("Please implement DisplayBattery() method")
}

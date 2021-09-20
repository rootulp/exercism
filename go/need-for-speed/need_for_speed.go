package speed

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}
}

// Track implements a race track.
type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery-car.batteryDrain < 0 {
		return car
	}
	nextBattery := car.battery - car.batteryDrain
	nextDistance := car.distance + car.speed
	return Car{
		speed:        car.speed,
		batteryDrain: car.batteryDrain,
		battery:      nextBattery,
		distance:     nextDistance,
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	numDrives := car.battery / car.batteryDrain
	possibleDistance := numDrives * car.speed
	return car.distance+possibleDistance >= track.distance
}

package lasagna

const MINUTES_PER_LAYER = 2

func OvenTime() int {
	return 40
}

func RemainingOvenTime(timeInOven int) int {
	return OvenTime() - timeInOven
}

func PreparationTime(layers int) int {
	return layers * MINUTES_PER_LAYER
}

func ElapsedTime(layers int, timeInOven int) int {
	return PreparationTime(layers) + timeInOven
}

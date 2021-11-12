package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) (sum int) {
	for _, birds := range birdsPerDay {
		sum += birds
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) (birdsInWeek int) {
	for i, birds := range birdsPerDay {
		if i/7 == week-1 {
			birdsInWeek += birds
		}
	}
	return birdsInWeek
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for day, birds := range birdsPerDay {
		if day%2 == 0 {
			birdsPerDay[day] = birds + 1
		}
	}
	return birdsPerDay
}

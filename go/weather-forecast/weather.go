// Package weather contains utilities for logging weath conditions
package weather

var (
	// CurrentCondition describes the current weather condition (e.g. "rainy")
	CurrentCondition string
	// CurrentLocation describes the current city (e.g. "New York")
	CurrentLocation string
)

// Log returns a string that describes the weath conditions in a particular city
func Log(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

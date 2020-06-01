package space

// Planet names
type Planet string

// Age returns the age in seconds of someone on a different planet.
func Age(ageInSeconds float64, planet Planet) float64 {
	if planet == "Earth" {
		return ageInSeconds / 31557600
	}
	return 0
}

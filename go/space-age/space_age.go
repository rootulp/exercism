package space

// Planet names
type Planet string

// Age returns the age in seconds of someone on a different planet.
func Age(ageInSeconds float64, planet Planet) float64 {
	if planet == "Earth" {
		return ageInSeconds / 31557600
	} else if planet == "Mercury" {
		return ageInSeconds / 31557600 / 0.2408467
	} else if planet == "Venus" {
		return ageInSeconds / 31557600 / 0.61519726
	}
	return 0
}

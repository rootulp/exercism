package space

// Planet names
type Planet string

// Age returns the age in years of someone on planet given their age in seconds.
func Age(ageInSeconds float64, planet Planet) float64 {
	if planet == "Earth" {
		return ageOnEarth(ageInSeconds)
	} else if planet == "Mercury" {
		return ageOnEarth(ageInSeconds) / 0.2408467
	} else if planet == "Venus" {
		return ageOnEarth(ageInSeconds) / 0.61519726
	}
	return 0
}

// AgeOnEarth returns the age of somone on Earth given their age in seconds.
func ageOnEarth(ageInSeconds float64) float64 {
	return ageInSeconds / 31557600
}

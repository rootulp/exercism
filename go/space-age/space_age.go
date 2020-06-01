package space

// Planet names
type Planet string

const secondsInEarthYear = 31557600

var orbitalPeriods = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age returns the age in years of someone on planet given their age in seconds.
func Age(ageInSeconds float64, planet Planet) float64 {
	if planet == "Earth" {
		return ageOnEarth(ageInSeconds)
	}
	return ageOnEarth(ageInSeconds) / orbitalPeriods[planet]
}

// AgeOnEarth returns the age of somone on Earth given their age in seconds.
func ageOnEarth(ageInSeconds float64) float64 {
	return ageInSeconds / secondsInEarthYear
}

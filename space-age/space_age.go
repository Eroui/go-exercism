package space

// Planet type definition
type Planet string

// Age calculates the age of a person living for seconds period in a given planet
func Age(seconds float64, planet Planet) float64 {
	secondsAYearInEarth := 31557600.0
	age := seconds / secondsAYearInEarth

	factors := map[Planet]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	return age / factors[planet]
}

// Package space provides Space Age solution.
package space

// yearInSeconds is seconds in a year.
const yearInSeconds = 31557600

// Planet represents a planet.
type Planet string

// ratios contains orbital period ratio of each planet to Earth years.
var ratios = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age converts seconds to years in given planet.
func Age(seconds float64, planet Planet) float64 {
	years := seconds / yearInSeconds
	return years / ratios[planet]
}

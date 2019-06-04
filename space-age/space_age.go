package space

import "math"

// Planet is the name of the planet
type Planet string

var orbitalPeriodMap = map[Planet]float64{
	"Mars":    1.8808158,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
	"Earth": 1,
}

const earthSeconds float64 = 31557600

// Age returns someones age on a planet
func Age(seconds float64, planet Planet) float64 {
	planetOrbitalPeriod := orbitalPeriodMap[planet]
	planetOrbitalPeriodSecs := planetOrbitalPeriod * earthSeconds
	ageOnPlanet := seconds / planetOrbitalPeriodSecs
	ageOnPlanetRounded := math.Round(ageOnPlanet*100) / 100
	return ageOnPlanetRounded
}

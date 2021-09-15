// Package weather provides tools to log weather conditions in a location.
package weather

var (
	// CurrentCondition represents the current weather condition.
	CurrentCondition string

	// CurrentLocation represents the current location.
	CurrentLocation string
)

// Log returns the weather log for the given city and condition.
func Log(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

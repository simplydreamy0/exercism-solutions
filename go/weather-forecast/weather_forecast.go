// Package weather blabla.
package weather

// CurrentCondition blabla.
var CurrentCondition string
// CurrentLocation blabla.
var CurrentLocation string

// Forecast blabla.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

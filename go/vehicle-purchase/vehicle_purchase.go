// Package purchase provides Vehicle Purchase solution.
package purchase

import "fmt"

// NeedsLicense determines whether a license is need to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(first, second string) string {
	if second < first {
		first = second
	}
	return fmt.Sprintf("%s is clearly the better choice.", first)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var price float64
	switch {
	case age >= 10:
		price = float64(originalPrice) * 0.5
	case age >= 3:
		price = float64(originalPrice) * 0.7
	default:
		price = float64(originalPrice) * 0.8
	}
	return price
}

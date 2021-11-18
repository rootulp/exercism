package purchase

import (
	"fmt"
	"strings"
)

// NeedsLicense determines whether a license is need to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	if strings.Compare(option1, option2) == 1 {
		return fmt.Sprintf("%s is clearly the better choice.", option2)
	}
	return fmt.Sprintf("%s is clearly the better choice.", option1)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * .8
	} else if age < 10 {
		return originalPrice * .7
	} else {
		return originalPrice * .5
	}
}

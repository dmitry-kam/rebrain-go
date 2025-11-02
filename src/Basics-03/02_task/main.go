package main

import (
	"fmt"
	"math"
)

func main() {
	type AmericanVelocity float32
	type EuropeanVelocity float32
	const kmInMile float32 = 1.609344

	var mPerSecond float32

	mPerSecond = 120.4
	kmPerHour := EuropeanVelocity(mPerSecond * 3600 / 1000)
	fmt.Printf("%.2f m/s = %.4f km/h\n", mPerSecond, kmPerHour)

	mPerSecond = 130.0
	milesPerHour := AmericanVelocity(mPerSecond * 3600 / 1000 / kmInMile)
	fmt.Printf("%.2f m/s = %.4f mph\n", mPerSecond, milesPerHour)

	roundedKmh := math.Round(float64(kmPerHour*100)) / 100
	roundedMph := math.Round(float64(milesPerHour*100)) / 100

	fmt.Printf("Rounded: %f km/h = %f (%.2f) mph\n", kmPerHour, roundedKmh, roundedKmh)
	fmt.Printf("Rounded: %f mph = %f (%.2f) mph\n", milesPerHour, roundedMph, roundedMph)
}

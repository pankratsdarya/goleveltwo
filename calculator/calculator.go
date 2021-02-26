// Package calculator implements some simple calculation methods.
// No validation - that was the task.
//
// The Borderlen() returns border length of circle on square given
//
// Borderlen(float64) float64
//
// The Diameter() returns diameter of circle on square given
//
// Diameter(float64) float64
//
// The Didgits() returns hundreds, tens and units in number given
//
// Didgits(uint32)
package calculator

import (
	"fmt"
	"math"
)

// Borderlen returns border length of circle
func Borderlen(s float64) float64 {
	return math.Sqrt(s/math.Pi) * 2 * math.Pi
}

// Diameter returns diameter of circle
func Diameter(s float64) float64 {
	return 2 * math.Sqrt(s/math.Pi)
}

// Didgits returns hundreds, tens and units in number
func Didgits(num uint32) {
	var hundreds, tens, units uint8
	hundreds = uint8(math.Floor(float64(num) / 100))
	num -= uint32(hundreds) * 100
	tens = uint8(math.Floor(float64(num) / 10))
	units = uint8(num % 10)
	fmt.Printf("В этом числе:\n сотен: %d,\n десятков: %d,\n единиц: %d. \n", hundreds, tens, units)
}

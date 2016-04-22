package goutils

import (
	"math"
)

func Sign(value float64) float64 {
	if value > 0 {
		return 1
	} else if value < 0 {
		return -1
	} else {
		return 0
	}
}

func Abs(value float64) float64 {
	return value * Sign(value)
}

func IsNegative(value float64) bool {
	return value < 0
}

func IsNonNegative(value float64) bool {
	return value > 0
}

func IsPositive(value float64) bool {
	return value > 0
}

func IsNonPositive(value float64) bool {
	return value <= 0
}

func IsWhole(value float64) bool {
	return Abs(math.Remainder(value, 1)) == 0
}

func IsNatural(value float64) bool {
	return IsWhole(value) && IsPositive(value)
}

func InRange(value, left, right float64) bool {
	if left > right {
		left, right = right, left
	}
	return value >= left && value <= right
}

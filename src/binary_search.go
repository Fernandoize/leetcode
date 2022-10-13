package main

import "math"

func Sqrt(n float64) float64 {
	low := 0.0
	high := n
	var mid = 0.0
	precision := 0.001
	for low <= high {
		mid = (low + high) / 2
		gap := math.Abs(mid*mid - n)
		if gap < precision {
			return mid
		}
		if mid*mid > n {
			high = mid - precision
		} else {
			low = mid + precision
		}
	}
	return mid
}

package main

import (
	"fmt"
	"math"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func main() {
	fmt.Println(toFixed(1.2345678, 0))  // 1
	fmt.Println(toFixed(19.2345678, 2)) // 1.2
	fmt.Println(toFixed(13.50000, 2))   // 1.23
}

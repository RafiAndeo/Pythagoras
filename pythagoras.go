package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	samping := pythagorasA(b, c)
	depan := pythagorasB(a, c)
	miring := pythagorasC(a, b)
	fmt.Println(samping, depan, miring)
}

func pythagorasA(b, c float64) float64 {
	a := (c * c) - (b * b)
	hasil := math.Pow(a, 0.5)
	return hasil
}

func pythagorasB(a, c float64) float64 {
	b := (c * c) - (a * a)
	hasil := math.Pow(b, 0.5)
	return hasil
}

func pythagorasC(a, b float64) float64 {
	c := (a * a) + (b * b)
	hasil := math.Pow(c, 0.5)
	return hasil
}

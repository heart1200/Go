package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (myFloat MyFloat) abs() float64 {
	if myFloat < 0 {
		return -float64(myFloat)
	}
	return float64(myFloat)
}

func main() {
	f := MyFloat(-math.SqrtPhi)
	fmt.Println(f.abs())
	fmt.Println(math.Sqrt(3), math.Sqrt2, math.SqrtE, math.SqrtPhi, math.SqrtPi)
}

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// func Abs(v Vertex) float64 {
// 	return math.Sqrt(v.x*v.x + v.y*v.y)
// }

func main() {
	v := Vertex{3, 4}
	// v1 := Abs(v)
	fmt.Println(v.abs())
	// fmt.Printf("v1 = %+v\n", v1)
}

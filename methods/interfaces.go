package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	var a Abser
	v := Vertex{3, 4}
	f := MyFloat(-math.Sqrt2)

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	fmt.Println(a.Abs())
}

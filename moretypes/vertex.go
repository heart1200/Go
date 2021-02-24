package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{1, 4}
	p := &v
	fmt.Println(*p)
	p.X = 1e9
	fmt.Println(v)
}

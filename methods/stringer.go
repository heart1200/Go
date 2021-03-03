package main

import "fmt"

type Persion struct {
	Name string
	Age  int
}

// Stringer是一个可以用字符串描述自己的类型。fmt包（还有很多包）都通过此接口来打印值。
func (p Persion) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Persion{"luodayu", 43}
	b := Persion{"Liyange", 41}
	fmt.Print(a, b)
}

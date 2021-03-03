package main

import "fmt"

func main() {
	var i interface{}
	i = "test"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	// 若 i 保存了一个 T，那么 t 将会是其底层值，而 ok 为 true。
	// 否则，ok 将为 false 而 t 将为 T 类型的零值，程序并不会产生恐慌。
	f, ok := i.(float32)
	fmt.Println(f, ok)

	f = i.(float32)
	fmt.Println(f)
}

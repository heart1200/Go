package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

// 即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用
// 保存了 nil 具体值的接口其自身并不为 nil
func main() {
	var i I

	var t *T
	i = t
	// fmt.Println(i)
	describe(i)
	i.M()

	i = &T{"test"}
	// fmt.Println("\n", *&i, "--------------")
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T\n)", i, i)
}

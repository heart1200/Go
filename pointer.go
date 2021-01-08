package main

//This is pointer
import "fmt"

func main() {
	i := 2
	p := &i         // & 是寻址符，指向变量i的地址，p就是指针
	fmt.Println(*p) // * 操作符表示指针指向的底层值

	*p = 4 // 通过指针修改底层值
	fmt.Println(i)
}

package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "array"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [7]int{1, 2, 3, 4, 5}
	fmt.Println(primes)

	primes1 := primes[1:4]
	fmt.Println(primes1)

	// 切片并不存储任何数据，它只是描述了底层数组中的一段。
	// 更改切片的元素会修改其底层数组中对应的元素。
	// 与它共享底层数组的切片都会观测到这些修改。
	primes1[1] = 66
	fmt.Println(primes)

	// make是专门用来创建slice,map,channel的值的，返回的是被创建的值，并且可以立即使用；
	// new是申请一小块内存并标记它是用来存放某个值的，返回的是指向这块内存的指针，而且这块内存不会被初始化
	b := make([]int, 5)
	printSlice("b", b)

	c := make([]int, 0, 5)
	printSlice("c", c)

	// 切片d是指向c的数组，c的长度是0，容量是5. 所以，d的长度是切了c的前2个值，容量还是5
	d := c[:2]
	printSlice("d", d)

	// e是指向d的切片，d的长度是2，容量是5。e从索引2开始取到5，所以，长度是3，容量是从2-5，为3
	e := d[2:5]
	printSlice("e", e)

	// 切片的增加
	var f []int
	var g []int
	printSlice1(f)

	f = append(f, 0)
	printSlice1(f)

	f = append(f, 1)
	printSlice1(f)

	f = append(f, 1, 2, 3)
	printSlice1(f)

	g = append(f, 5, 6)
	printSlice1(g)

}
func printSlice(s string, x []int) {
	fmt.Printf("%s, len=%d, cap=%d, %v\n", s, len(x), cap(x), x)
}

func printSlice1(x []int) {
	fmt.Printf("x, len=%d, cap=%d, %v\n", len(x), cap(x), x)
}

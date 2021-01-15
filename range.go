package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func main() {
	// for循环可以遍历切片和映射，使用range返回2个值，第一个是下标，第二个是返回对应下标的副本值
	var s = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range s {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for i, v := range pow {
		fmt.Printf("%d, %d\n", i, v)
	}

	pic.Show(drawPic)

}

func drawPic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for y := range a {
		b := make([]uint8, dx)
		for x := range b {
			b[x] = uint8(x ^ y)
		}
		a[y] = b
	}
	return a
}

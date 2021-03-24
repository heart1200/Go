package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

//和映射与切片一样，信道在使用前必须创建
func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// x, y := <-c, <-c
	// fmt.Println(x, y, x+y)
	fmt.Print(<-c, <-c)
}

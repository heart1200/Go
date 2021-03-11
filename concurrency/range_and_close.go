package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < 10; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

//循环 for i := range c 会不断从信道接收值，直到它被关闭
func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Printf("%v  ", i)
	}
}

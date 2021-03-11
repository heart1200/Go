package main

import (
	"fmt"
	"time"
)

func TestGoroutine(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go TestGoroutine("Dayu")
	TestGoroutine("Winlin")
}

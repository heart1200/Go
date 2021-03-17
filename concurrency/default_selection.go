package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	after := time.Tick(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-after:
			fmt.Println("Booms")
			return
		// 当 select 中的其它分支都没有准备好时，default 分支就会执行
		default:
			fmt.Println("    :")
			time.Sleep(50 * time.Millisecond)
		}

	}
}

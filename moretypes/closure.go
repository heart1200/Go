package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func(int) int {
	fib0, fib1 := 0, 1
	return func(int) int {
		fib0, fib1 = fib1, (fib0 + fib1)
		return fib0
	}
}

func main() {
	// pos, mov := adder(), adder()
	f := fibonacci()
	for i := 0; i < 10; i++ {
		// fmt.Println(pos(i), mov(-2*i+10))
		fmt.Println(f(i))
	}

}

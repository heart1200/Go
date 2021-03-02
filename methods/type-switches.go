package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v twice is %v\n", v, 2*v)
	case string:
		fmt.Printf("%q is %v byte\n", i, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(5)
	do("Luodayu")
	do(true)
}

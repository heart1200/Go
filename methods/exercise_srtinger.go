package main

import "fmt"

type IPAdd [4]byte

func (i IPAdd) IpFormat() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

func main() {
	hosts := map[string]IPAdd{
		"tiantan": {10, 100, 201, 118},
		"shunyi":  {192, 168, 1, 124},
	}
	for name, ip := range hosts {
		fmt.Printf("%v：%v\n", name, ip)
	}
}

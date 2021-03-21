// 明明想在学校中请一些同学一起做一项问卷调查，为了实验的客观性，他先用计算机生成了N个1到100
// 0之间的随机整数（N≤1000），对于其中重复的数字，只保留一个，把其余相同的数去掉，不同的数对
// 应着不同的学生的学号。然后再把这些数从小到大排序，按照排好的顺序去找同学做调查。请你协助明
// 明完成“去重”与“排序”的工作(同一个测试用例里可能会有多组数据(用于不同的调查)，希望大家能正
// 确处理)。

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func random_id(n, source int) []int {
	rand.Seed(time.Now().UnixNano())
	var i int
	if n%2 == 1 {
		i = 1
	}
	id_1 := rand.Perm(source)[0 : n/2+i]
	id_2 := rand.Perm(source)[0 : n/2]
	id := append(id_1, id_2...)
	return id
}

// func sort_for_random(a []int) []int {
// 	return sort.Ints(a)
// }

func unrepeated(a []int) []int {
	i := 0
	for j := 1; j < len(a); j++ {
		if a[i] != a[j] {
			i++
			a[i] = a[j]
		}
	}
	return a[:i+1]
}

func main() {
	// fmt.Println(random_id(10, 15))
	// a := []int{1, 1, 1, 2, 2, 3, 1}
	// fmt.Println(unrepeated(a))
	a := random_id(300, 1000)
	sort.Ints(a)
	new := unrepeated(a)
	fmt.Printf("The id is: %v\n the length is: %v\n", new, len(new))
	// fmt.Println(unrepeated(new_a))
}

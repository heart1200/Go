package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	fmt.Println(t.Value)
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	m := make(map[int]int)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()
	for {
		//将ch1的数据存入map
		v, f := <-ch1
		if f {
			k, _ := m[v]
			m[v] = k + 1
		} else {
			break
		}
	}
	for {
		//比较ch1的数据
		v, f := <-ch2
		if f {
			k, _ := m[v]
			if k == 0 {
				return false
			} else {
				m[v] = k - 1
			}
		} else {
			break
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(2), tree.New(2)))
}

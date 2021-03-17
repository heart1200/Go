// 题目：计算字符串最后一个单词的长度，单词以空格隔开。
// 输入一行，代表要计算的字符串，非空，长度小于5000。

package main

import (
	"fmt"
	"strings"
)

// 使用函数库来解决
func count_last_word_in_func(s string) {

	if len(s) > 0 && len(s) < 5000 {
		strs := strings.Split(strings.TrimSpace(s), " ")
		// fmt.Print(len(strs[len(strs)-1]))
		fmt.Printf("The last word length is: %v\n", len(strs[len(strs)-1]))
	} else {
		fmt.Println("The string is too much or null")
	}
}

// 使用常规方法解决,从后向前，直到遇到第一个空格
func count_last_word_in_normal(s string) int {
	length := 0
	byteS := []byte(s)
	// fmt.Print(byteS)
	if len(s) > 0 && len(s) < 5000 {
		for i := len(s) - 1; i >= 0; i-- {
			if byteS[i] == 32 {
				// fmt.Println(i)
				break
			}
			length++
		}
	} else {
		fmt.Println("The string is too long or null")
	}
	return length
}

func main() {
	count_last_word_in_func("hello word luodayu, and luoweilin")
	fmt.Println(count_last_word_in_normal(""))
	//fmt.Println(lengthOfLastWord("hello luowelin"))
}

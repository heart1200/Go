// 写出一个程序，接受一个由字母、数字和空格组成的字符串，和一个字母，然后输出输入字符串中该字母的出现次数。不区分大小写。

package main

import (
	"fmt"
	"strings"
)

// 利用循环进行统计
func count_letter_in_for(sentence, letter string) int {
	var count = 0
	for i, j := 0, 1; j < len(sentence); {
		if strings.EqualFold(letter, sentence[i:j]) {
			count++
		}
		i = i + 1
		j = j + 1
	}
	return count
}

// 利用函数
func count_letter_in_func(sentence, letter string) int {
	count_lower := strings.Count(sentence, strings.ToLower(letter))
	count_upper := strings.Count(sentence, strings.ToUpper(letter))
	return count_lower + count_upper
}

func main() {
	fmt.Println(count_letter_in_for("ABCabc", "A"))
	fmt.Println(count_letter_in_func("ABCabc", "A"))
	fmt.Println(strings.SplitN("heart", "", 3))
}

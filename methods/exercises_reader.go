package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// 实现一个 Reader 类型，它产生一个 ASCII 字符 'A' 的无限流
func (m MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}

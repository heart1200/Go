package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(b []byte) (int, error) {
	n, err := reader.r.Read(b)
	for i := 0; i < n; i++ {
		switch {
		case b[i] >= 'A' && b[i] < 'N':
			b[i] += 13
		case b[i] >= 'N' && b[i] <= 'Z':
			b[i] -= 13
		case b[i] >= 'a' && b[i] < 'n':
			b[i] += 13
		case b[i] >= 'n' && b[i] <= 'z':
			b[i] -= 13
		}
	}
	return n, err

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	// fmt.Printf("s = %v\n", *s)
	read := rot13Reader{s}
	// fmt.Printf("read = %v\n", read)
	// fmt.Printf("read.r = %v\n", read.r)
	io.Copy(os.Stdout, &read)
}

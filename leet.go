package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func scan(a ...interface{})             { fmt.Fscan(reader, a...) }
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

func main() {
	s := []byte("hello")
	fmt.Println(string(s))
	right := len(s) - 1
	left := 0
	for right > left {
		temp := s[right]
		s[right] = s[left]
		s[left] = temp
		right--
		left++

	}
	fmt.Println(string(s))
}

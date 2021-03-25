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

//https://codeforces.com/problemset/problem/1499/B
func main() {
	defer writer.Flush()
	var n int
	foundZero := false
	scan(&n)
	a := make([]string, n)
	for i := 0; i < n; i++ {
		scan(&a[i])
	}
	for i := 0; i < n; i++ {
		if foundZero {

		}
	}

}

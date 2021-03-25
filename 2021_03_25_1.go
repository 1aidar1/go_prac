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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//https://codeforces.com/problemset/problem/1501/B
func main() {
	defer writer.Flush()
	var t, n int
	scan(&t)
	out := make([][]int, t)

	for i := 0; i < t; i++ {
		scan(&n)
		a := make([]int, n)
		for j := 0; j < n; j++ {
			scan(&a[j])
		}
		ans := make([]int, n)
		med := 0
		for j := n - 1; j >= 0; j-- {
			med = max(med, a[j])
			if med > 0 {
				ans[j] = 1
			} else {
				ans[j] = 0
			}
			med--
		}
		out[i] = ans
	}
	for i := 0; i < t; i++ {
		for j := 0; j < len(out[i]); j++ {
			fmt.Printf("%d ", out[i][j])
		}
		fmt.Println()
	}
}

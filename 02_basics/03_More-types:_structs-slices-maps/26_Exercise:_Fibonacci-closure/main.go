package main

import (
	"fmt"
)

// fibonacci is a function that returns a function that returns an int.
func fibonacci() func() int {
	n0, n1 := 0, 1 // 変数に代入する一番最初だけ呼び出される
	return func() int {
		v := n0
		n0, n1 = n1, n0+n1
		return v
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

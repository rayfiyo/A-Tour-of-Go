// スライスへの append （付加）
// Slices: usage and internals https://go.dev/blog/slices-intro

package main

import (
	"fmt"
)

func main() {
	var s []int
	printSlice(s)

	// nil スライスにも append は機能する
	// append works on nil slices
	s = append(s, 0)
	printSlice(s)

	// スライスは必要に応じて拡大する
	// The slice grows as needed
	s = append(s, 1)
	printSlice(s)

	// 一度に複数の要素を追加できる
	// We can add more than one element at a time
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

// 型を記述する位置
// https://go.dev/blog/declaration-syntax

package main

import "fmt"

// 戻り値に変数名をつけることができる（named return value）
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // "naked" return という
}

func main() {
	fmt.Println(split(17))
}

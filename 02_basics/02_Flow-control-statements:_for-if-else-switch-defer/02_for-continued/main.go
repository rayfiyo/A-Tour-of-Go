package main

import "fmt"

func main() {
	// ステートメントは省略できる
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

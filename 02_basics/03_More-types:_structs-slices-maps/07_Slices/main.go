// スライス（可変長配列）の利用
// 2, 3, 5, 7, 11, 13 の数列 primes

package main

import (
	"fmt"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // 半開区間で表現，明示的に非推奨の記述
	// # 半開区間のメリット
	// - slice[a:a] で空の配列の表現
	// - slice[a:b] だと b-a で要素数がわかる

	fmt.Println(s)
}

// スライスの長さと容量
// スライスは配列のポインタ的なものなので，長さを換えても元の配列は変更されない

package main

import (
	"fmt"
)

func main() {
	// 容量は len([]int{2, 3, 5, 7, 11, 13}) となる
	s := []int{2, 3, 5, 7, 11, 13}
	printSclice(s)

	// 長さ０にスライス(slice)
	s = s[:0]
	printSclice(s)

	// スライスはポインタ的なものなので，スライスしても元の配列は変更されない
	// 実際に次のように伸ばすことができる

	// 長さ４に伸ばす(extent)
	s = s[:4]
	printSclice(s)

	// （元の配列 []int{2, 3, 5, 7, 11, 13} から）
	// 先頭より２文字削る(drop)
	s = s[2:]
	printSclice(s)
}

func printSclice(s []int) {
	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)
}

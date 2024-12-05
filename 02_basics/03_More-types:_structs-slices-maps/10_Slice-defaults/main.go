// スライスするときは既定値を使うことで略記できる
// 既定値は，下限: ０，上限: スライスの長さ

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	// 通常にスライスする
	s = s[1:4]
	fmt.Println(s)

	// 下限を省略
	s = s[:2]
	fmt.Println(s)

	// 上限を省略
	s = s[1:]
	fmt.Println(s)
}

/*
また，
var a [10]int
において，以下のスライス式は等価
a[0:10]
a[:10]
a[0:]
a[:]
*/

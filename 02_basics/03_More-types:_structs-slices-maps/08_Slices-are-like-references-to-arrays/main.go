// スライスは配列のポインタのようなもの
// 	"John",	"Paul",	"George", "Ringo" の文字列配列 names

package main

import (
	"fmt"
)

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	// 一言で言えば names[1] = "XXX" と等価な操作
	// スライスの要素を変更すると，元となる配列に対応する要素が変更される
	// スライスは配列のポインタのようなもので，
	// どんなデータも格納しておらず，単に元の配列の部分列を指し示しているだけ

	fmt.Println(a, b)
	fmt.Println(names)
}

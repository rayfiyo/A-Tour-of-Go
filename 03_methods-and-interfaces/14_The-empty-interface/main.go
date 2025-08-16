// 空のインターフェース: 0 個のメソッドを指定したインターフェース型 interface{}
//     任意の型の値を保持できる（すべての型は、0個以上のメソッドを実装しているから）
//     未知の型の値を扱うときに使う
//     Go 1.8+ では any エイリアスが利用できる

// 主な処理: i 抽象型を定義し、describe関数を実行し、
//     続いて i = 42, i = "hello" の場合の describe をそれぞれ実行する
// describe関数: i 抽象型を引数にとり、その値と型を標準出力する

package main

import (
	"fmt"
)

func main() {
	var i any
	// var i interface{}

	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i any) {
	// func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

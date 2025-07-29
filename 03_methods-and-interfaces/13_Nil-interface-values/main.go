// インターフェースが（完全に） nil のとき、動的値も動的型も保持しない
// このとき、メソッド呼び出し（例: i.M()）は、
// 呼び出された時点でパニック（ランタイムエラー）になる
//
// よって、linter が警告を出す
// Diagnostics:
// 1. nil dereference in dynamic method call [nilderef]

// I 抽象型: メソッド M の集合
// 主な処理:
//     1. i I型を定義し、describe()し、メソッドMを実行する
// describe関数: i I型を引数にとり、その値と型を標準出力する

package main

import (
	"fmt"
)

type I interface {
	M()
}

// M の実装を行っても結果は変わらない

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

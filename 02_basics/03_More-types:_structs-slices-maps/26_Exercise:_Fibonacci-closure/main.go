package main

import (
	"fmt"
)

// fibonacci is a function that returns a function that returns an int.
// フィボナッチは int を返す関数を返す関数
func fibonacci() func() int {
	// 変数に代入する一番最初だけ呼び出される
	// "fibonacci関数スコープ" の変数が定義とも言える
	n0, n1 := 0, 1

	return func() int {
		// 戻り値に使う値を一時変数vに保持
		v := n0

		// 次回向けに状態を更新
		// n1 は次のフィボナッチ数が格納されているので n0 に代入
		// 次の次（n1の次）の値を計算したものを n1 に代入
		// n0 = n1 ; n1 = n0 + n1 とは違う．こちらは２の累乗になる
		n0, n1 = n1, n0+n1

		return v
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

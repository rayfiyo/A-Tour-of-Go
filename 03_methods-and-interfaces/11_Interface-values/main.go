// インターフェース（の値）は、動的型 と 動的値 のタプルである。と説明できる
// 例: `var i interface{}` は、どちらも nil
// 例: `var p *T = nil; var j interface{} = p` は、動的型は *T で、動的値が nil

// I 抽象型: M をレシーバに持つ
// T 型: string 型のフィールド S を持つ
// *T 型のレシーバ引数 t に関する M の実装: t.S を標準出力する
// F 型: float64
// *F 型のレシーバ引数 f に関する M の実装: f を標準出力する
// 主な処理: T 型を describe して M を実行、 F 型も同じ
// describe 関数: I 型の引数 i の値（値と型）を標準出力
//				  fmt.Printf("(%v, %T)\n", i, i)

package main

import (
	"fmt"
	"math"
)

type I interface {
	// どんな型でもいいけど、M() は実装しておいてね～
	M()
}

type T struct {
	S string
}

// T 型で M() を実装しているので、T 型は I インターフェースを満たす
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

// F 型で M() を実装しているので、F 型は I インターフェースを満たす
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i) // インターフェースを満たしているのでエラーにならない
	i.M()

	i = F(math.Pi)
	describe(i) // インターフェースを満たしているので（略）
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

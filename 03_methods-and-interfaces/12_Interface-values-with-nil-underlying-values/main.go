// インターフェースそのもの（動的値）が nil の場合、レシーバは nil をとる
// 他の言語と比較すると、Go は null ポインター例外 にはならない
// だから、そのインターフェースが nil を許容する場合は if 文を使って処理する
//
// インターフェース t の動的型が nil を値に持つポインタの場合、
// 動的値はポインタが示す nil を返すことに注意が必要である
// つまり、動的値が nil であっても、動的型は 非nil であるから `t != nil` となる
// 動的値を取り出したい場合は型アサーションを使う必要がある
// if ptr, ok := t.(*T); ok {
//     if ptr == nil {
// 	   }
// }

// I 抽象型: メソッド M の集合
// T 構造体: S 文字列型をフィールドにもつ
// メソッドMの実装: t *T型をレシーバに持つ
//     I 抽象型は nil を許容する
//     t.S の値を標準出力する
// 主な処理:
//     1. i I型に、*T 型の nil 値 t を代入し、describe()し、メソッドMを実行する
//     2. i I型に、*T 型の "hello" 値を代入し、 1 と同様に処理する
// describe関数: i I型を引数にとり、その値と型を標準出力する

package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T // 型付き nil はここから来た（別の例: t := hoge()）
	i = t
	// 上は `var i I = (*T)(nil)` とも表現できるが、
	// 中間変数 t を用いた方が どこから型付き nil が来たかわかりやすい
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

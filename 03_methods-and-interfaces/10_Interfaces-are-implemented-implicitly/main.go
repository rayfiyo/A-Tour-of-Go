// ある型のメソッドを実装するとは、その型をレシーバに持つ関数を定義するということ。
// インターフェースのフィールド？に、あるメソッドのシグネチャを記述することでメソッドを関連付ける？
// そして、そのメソッドを実装することで、インターフェースを実装する（満たせる）。

// I interface: メソッド M の集合
// T 型: string の文字列の変数 S をフィールドに持つ
// t.S を返すメソッド M:（T 型のメソッドを実装）、 レシーバ引数に T 型の値 t を持つ

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

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()
}

// ポインタレシーバのメソッドの利点として、（ポインタ引数の関数と違って）
// 呼び出す時の型がポインタかどうかを意識しなくてよい。
// （アドレス可能な値に対して自動でアドレスが取られるので、
//   呼び出し時に型を意識する必要がない）

// Vertex 型: float64 の２つのフィールド X, Y を持つ
// Scale メソッド: Vertex 型の値 v を float64 の値 f 倍する
// ScaleFunc 関数: Scale メソッドと同じ

package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

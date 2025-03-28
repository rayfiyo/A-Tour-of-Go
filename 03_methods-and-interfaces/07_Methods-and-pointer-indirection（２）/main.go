// 前回、Go はレシーバをポインタで持つメソッドの場合、
// 関数内部でアドレス可能な値に対して自動でアドレスが取られると述べた。
// 一方、レシーバを値で持つメソッドの場合でも 似たような関係として、
// ポインタで呼び出す時に、ポインタから自動的に参照解除され、値として呼び出せる。

// Vertex 型: float64 の２つのフィールド X, Y を持つ
// Abs メソッド: Vertex 型の値 v をレシーバ引数にとり、２値のニ乗和の平方根を返す
// AbsFunc 関数: Abs メソッドと同じ

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs()) // (*p).Abs() として解釈
	fmt.Println(AbsFunc(*p))
}

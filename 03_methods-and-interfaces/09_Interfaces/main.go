// はじめての Interface
// メソッドのシグニチャの集まりで定義する
// そのメソッドを持つ型ならどれでもそのインターフェース型として扱えるようになる

// Abser interface: Abs メソッドを持つ
// MyFloat型: float64
// Absメソッド（メソッド引数: MyFloat型の値 f）: f が 0 より小さければ -1 倍して返す
// Vertex型: X と Y それぞれ float64
// Absメソッド（メソッド引数: *Vertex型の値 v）: v のフィールド2値のニ乗和の平方根を返す

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // MyFloat は Abser を実装している
	a = &v // *Vertex は Abser を実装している

	// v は Vertex型（*Vertex型ではない）で、実装していない
	// a = v

	fmt.Println(a.Abs())
}

// MyFloat の Abser 実装

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

// Vertex の Abser 実装

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

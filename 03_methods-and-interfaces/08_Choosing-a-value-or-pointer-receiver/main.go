// ポインタレシーバを使う理由は、
// １．メソッドでレシーバが指す変数を変更するため（メソッド外で影響）
// ２．（特に大きな構造体のレシーバで）メソッドの呼び出し毎に変数を複製させないため
// また、（１つのパッケージ内で？）値レシーバまたはポインタレシーバのどちらかに統一する

// Vertex 型: float64 の２つのフィールド X, Y を持つ
// Scale メソッド: Vertex 型の値 v を float64 の値 f 倍する
// Abs メソッド: Vertex 型の値 v をレシーバ引数にとり、２値のニ乗和の平方根を返す

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}

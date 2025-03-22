// レシーバをポインタで宣言できる
// 例えば、返り値がいらない場合にポインタレシーバをとる実装にする

// Vertex 型: float64 の２つのフィールド X, Y を持つ
// Abs メソッド: Vertex 型の値 v をレシーバ引数にとり、２値のニ乗和の平方根を返す
// Scale メソッド: Vertex 型の値 v を float64 の値 f 倍する

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// こっちは値を書き換えたいわけではなく、計算結果がほしい
// 普通の関数のような動作が望まれている
// だから return が必要、ポインタは不適切
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 値の変更がしたく、返り値はいらないのでポインタが望ましい
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

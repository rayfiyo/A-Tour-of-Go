// ポインタと値の違い

// Vertex 型: float64 の２つのフィールド X, Y を持つ
// Abs 関数: Vertex 型の値 v を引数にとり、２値のニ乗和の平方根を返す
// Scale 関数: Vertex 型の値のポインタ v を float64 の値 f 倍する

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}

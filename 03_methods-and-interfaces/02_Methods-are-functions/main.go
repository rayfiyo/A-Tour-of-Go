// 前回のメソッドの実装を関数で実装した

// Vertex 型: float64 の２つのフィールド X, Y を持つ
// Abs 関数: Vertex 型の値 v を引数にとり、２値のニ乗和の平方根を返す

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

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}

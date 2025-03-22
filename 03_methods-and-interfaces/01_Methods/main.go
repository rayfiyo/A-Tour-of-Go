// Goは型にメソッドを指定できる
// メソッド: 引数として、レシーバ引数を自身にとる関数
// レシーバ: func 句とメソッド名の間に自身の引数リストで表現

// Vertex 型: float64 の２つのフィールド X, Y を持つ
// Abs メソッド: Vertex 型の値 v をレシーバ引数にとり、２値のニ乗和の平方根を返す

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

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

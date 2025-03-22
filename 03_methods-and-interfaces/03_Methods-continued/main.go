// レシーバを伴うメソッドの宣言は、そのレシーバ型が同じパッケージにある必要がある
// 組み込みの型も別パッケージであることに注意する

// MyFloat 型: float64 のラップ
// Abs メソッド: MyFloat 型の値 f をレシーバ引数にとり、
//               負の値であれば正の float64 型に、そうでなければ float64 型にして返す

package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

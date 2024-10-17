// defer の実行
// defer は，通常通り評価（値の確定）がされるが，実行はその関数の return 直前まで保持される．
package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("world") // 評価はされるが実行はされない

	fmt.Println("hello")
	// （イメージ的な説明として）ここで defer が実行される．
}

// 評価について
/*
a := 3 + 2*4
は，右辺が評価されると 11 となる．その後，代入（実行の１つ）される．

b := 2
func(c)
を評価すると func(2) となる．
*/

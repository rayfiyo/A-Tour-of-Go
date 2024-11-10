package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // i のポインタを p として定義
	fmt.Println(*p) // i を出力（i のポインタ p から実体参照）
	*p = 21         // i に 21 を代入（i のポインタ p を操作）
	fmt.Println(i)  // i を出力

	p = &j         // j のポインタを p として定義
	*p = *p / 27   // i を 27 で割る（i のポインタ p を実体参照）
	fmt.Println(j) // i を出力
}

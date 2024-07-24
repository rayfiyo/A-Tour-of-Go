// const は値だから型を持たない（必要に応じて動的に定まる）．

package main

import "fmt"

const (
	// 1 ビットを 100 桁左にシフトして巨大な数を作成します．
	// 言い換えれば，1 の後に 100 個のゼロが続く 2 進数です．
	Big = 1 << 100
	// 再び 99 桁右にシフトすると，1<<1 または 2 になります．
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

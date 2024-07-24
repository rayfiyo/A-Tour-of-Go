// 定数 constant は const で宣言
// const は 文字，文字列，真理値，数値

package main

import "fmt"

// := ではない
const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rlues?", Truth)
}

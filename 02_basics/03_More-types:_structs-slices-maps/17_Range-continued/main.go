// `_` 代入で捨てることができる
// 非本質: `1 << B` で 1 を左に B ビットシフトできる
//          B は uint である必要があるが，int でも動作する（負の範囲ではパニック）
// 考察: uint の実態は int で，uint() は負の値だと 0 を返しているだけ？
// 検証: -> どちらも 8 byte だったが，uint のほうが大きな値を格納できる
//       つまり，前述の考察は正しくなさそう

package main

import (
	"fmt"
)

func main() {
	pow := make([]int, 10)
	for i := range pow {
		// ビットシフト演算を行う
		// つまり 2^i と等価
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

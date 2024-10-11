// switch の使い方
// C と違って break は必要ない（自動的に break）
// tips: time パッケージでの曜日は int 型

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday { // 適切な型（int）で土曜日を表現
	case today + 0: // 明示的に加算
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// スライスに他のスライスを含める（任意の型を含められる）

package main

import (
	"fmt"
	"strings"
)

func main() {
	// 五目並べの盤作成
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
		/* 明示すると以下
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		*/
	}

	// プレイ履歴
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	// 出力
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

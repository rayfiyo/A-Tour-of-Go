// リテラル: 即値，変数ではなく固定値，スクリプト内に文字で記述
// 例えば，[3]bool{true, true, false} は配列リテラルで，
// []bool{true, true, false} はそれと同じ配列を作成し参照するスライスのリテラル
// 2, 3, 5, 7, 11, 13 の配列 q []int
// true, false, true, true, false, true の配列 r []bool
// {2, true}, {3, false}, {5, true}, {7, true}, {11, false}, {13, true} の配列 s []struct { i int, b bool}

package main

import (
	"fmt"
)

func main() {
	// int のスライスリテラル
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	// bool のスライスリテラル
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// 構造体（int と bool） のスライスリテラル
	s := []struct {
		// フィールドに変数名がない場合は，アクセスするときに型を用いる
		// 例 s[6].i は s[6].int となる
		i int
		b bool
	}{
		// 特に法則性はない
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

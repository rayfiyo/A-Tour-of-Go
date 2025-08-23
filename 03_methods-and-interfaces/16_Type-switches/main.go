// 型スイッチ: 一連の型アサーションをまとめて行うための構文
//     通常の switch 文に似ている
//     違いとして、型スイッチにおける case では値ではなく型を指定する
//     （その指定した型と、与えられたインターフェースの動的型と比較される）
// 型スイッチ内での宣言: 型アサーションと同じ構文に、type キーワードを用いる
//     例: v := i.(type)
// 	       case の中では、変数 v は case で指定した型になり、i が保持する値を格納する
//         default の中では、v は i と同じインターフェース型および値を持つ

// do 関数: i という空のインターフェースを引数に持つ
//     型が int なら、fmt.Printf("Twice %v is %v\n", v, v*2)
//     型が string なら、fmt.Printf("%q is %v bytes long\n", v, len(v))
//     それ以外は、fmt.Printf("I don't know about type %T!\n", v)
// 主な処理: do 関数を、21, "hello", true それぞれで実行する

package main

import (
	"fmt"
)

func do(i any) {
	// func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}

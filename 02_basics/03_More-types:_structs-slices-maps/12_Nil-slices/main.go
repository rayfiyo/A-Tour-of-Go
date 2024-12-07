// スライスの初期値，nil

package main

import (
	"fmt"
)

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil { // tautological condintion
		fmt.Println("nil!")
	}
}

// Map の基礎

package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var (
	l map[string]Vertex
	m map[string]Vertex
)

func main() {
	if nil == m {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
		fmt.Println(l)
		fmt.Println(m)
	}

    // マップを初期化し，使用可能な状態にする
	m = make(map[string]Vertex)

	if nil == m {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
		fmt.Println(l)
		fmt.Println(m)
	}
}

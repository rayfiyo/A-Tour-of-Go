// 構造体のポインタ参照

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 // (*p).X と同じ意味
	// (*p).X = 1e3
	fmt.Println(v)
}

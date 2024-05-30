package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// rand.New(rand.NewSource(2)) // Seed を2に設定
	fmt.Println("My favorite number is", rand.Intn(10))
}

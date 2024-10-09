// switch 関数
// Goでは break が自動で適用され，適合する条件以降の case は実行されない
// if のようにステートメントの記述が可能
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows, ...
		fmt.Printf("%s.\n", os)
	}
}

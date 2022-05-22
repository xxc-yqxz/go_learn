package main

import "fmt"

func main() {
	var a [3]int
	a[0], a[1], a[2] = 1, 2, 3
	fmt.Println(a)
	var idx4 int = 4
	// var idx4 int = -1
	a[idx4] = 4 // 报错
	fmt.Println(a)
}

package main

import "fmt"

func main() {
	var a, b int8 = 30, 10
	fmt.Println("a * b = ", a*b) // 44 -> 可以看到与正确结果不符
	fmt.Println(30 / 40)         // 0
}

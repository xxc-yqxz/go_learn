package main

import "fmt"

func main() {
	const pi = 3.14159265
	fmt.Println(pi)
	// pi = 3.14	// 此处 pi 不可以再被赋值，因为常量一旦定义就不能再被修改了
	const a = 123456 // 常量定义不使用不会报错
}

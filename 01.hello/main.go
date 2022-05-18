package main

import "fmt"

func main() {
	var name string = "hello,golang!"
	fmt.Println(name)
	var age int = 35
	fmt.Println(age)
	var tall float64 = 1.70
	fmt.Println(tall)

	name = "哈哈哈"
	// age = "测试" // 类型不兼容报错
	// tall = age  // 类型不兼容报错
}

package main

import "fmt"

func main() {
	fmt.Println("中文，abc，123，❤")
	fmt.Printf("我的名字是%s\n", "小强") // 我的名字是小强
	fmt.Printf("我的名字是%q\n", "小强") // 我的名字是"小强"
	fmt.Printf("我的名字是%x\n", "小强") // 我的名字是e5b08fe5bcba	--> 转换为utf8字符串编码
	fmt.Printf("我的名字是%X", "小强")   // 我的名字是E5B08FE5BCBA
}

package main

import "fmt"

func main() {
	var money int // 修改它，看看它会做什么
	if money <= 20 {
		fmt.Println("点个外卖")
	} else if money <= 200 {
		fmt.Println("去下馆子")
	} else if money <= 2000 {
		fmt.Println("到米其林体验体验")
	} else if money <= 20000 {
		fmt.Println("出国游玩一圈")
	} else {
		fmt.Println("容我想想")
	}
}

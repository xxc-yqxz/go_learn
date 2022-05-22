package main

import "fmt"

func main() {
	var money = 0
	var busy bool = false
	switch {
	case money >= 0 && money <= 20:
		fmt.Println("点1个外卖")
		if busy {
			break
		}
		fmt.Println("再吃个零食")
		fallthrough // 直接并入下一个处理分支而无需判断条件
	case money > 20 && money <= 200:
		fmt.Println("去下馆子")
	default:
		fmt.Println("容我想想")
	}
	fmt.Println("结束")
}

package main

import "fmt"

func main() {
	a, b := 100, 31
	fmt.Println(a ^ b) // 123
	fmt.Println(b ^ a) // 123  -> 可以看到，^ 满足交换律

	// ^ 的应用：找到一个每个数字都出现两次中的多余数字
	arr := []int{4, 3, 4, 5, 6, 7, 3, 6, 5}
	result := -1
	for _, item := range arr {
		if result < 0 {
			result = item
		} else {
			result = result ^ item
		}
	}

	fmt.Println(result) // 7
}

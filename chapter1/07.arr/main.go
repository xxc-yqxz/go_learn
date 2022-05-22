package main

import "fmt"

func main() {
	var age int = 35
	fmt.Println(age)
	var ages [5]int = [5]int{35, 32, 33, 37, 59}
	var ages2 = [5]int{35, 32, 33, 37, 59}
	ages3 := [5]int{35, 32, 33, 37, 59}
	fmt.Println(ages, ages2, ages3)
	ages2 = ages3
	// ages3 = [6]int{35, 32, 33, 37, 59} // 错误：数组长度不能变

	ages4 := [...]int{1, 2, 3, 4} // 使用 [...] ，便可以让ages4的长度根据值进行改变
	fmt.Println(ages4)

	var ages5 [3]int
	fmt.Println(ages5) // [0 0 0]
	ages5[0] = 1000
	ages5[1] = 1111
	ages5[2] = 2222
	fmt.Println(ages5)
	//ages5[-1] = -1 // 错误，越界
	//ages5[99] = -1 // 错误，越界
	for i := 0; i < len(ages4); i++ {
		fmt.Println(ages4[i])
	}
	for i, val := range ages4 {
		fmt.Println(ages4[i], "===>", val, "===>", i)
	}
	// 1 ===> 1 ===> 0
	// 2 ===> 2 ===> 1
	// 3 ===> 3 ===> 2
	// 4 ===> 4 ===> 3
}

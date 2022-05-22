package main

import "fmt"

func main() {
	var money interface{} = 10
	// var money interface{} = 10.0
	// var money interface{} = "10"
	switch money.(type) {
	case int, int64, int16, int32:
		fmt.Println("money是 整数")
	case float64, float32:
		fmt.Println("money是 小数")
	default:
		fmt.Println("money是未知类型")
	}
	fmt.Println("结束", money)
}

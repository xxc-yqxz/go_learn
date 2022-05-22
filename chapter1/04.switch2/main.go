package main

import "fmt"

func main() {
	//var money interface{} = 10
	var money interface{} = 10.0
	// var money interface{} = "10"
	switch money.(type) {
	case int:
		fmt.Println("money是 int")
	case int64:
		fmt.Println("money是 int64")
	case float64:
		fmt.Println("money是 float64")
	case float32:
		fmt.Println("money是 float32")
	default:
		fmt.Println("money是未知类型")
	}
	// money = 3  // 此处不会报错
	fmt.Println("结束", money)
}

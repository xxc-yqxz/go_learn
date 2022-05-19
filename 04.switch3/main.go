package main

import "fmt"

func main() {
	//var money interface{} = 10
	var money interface{} = 10.0
	// var money interface{} = "10"
	switch newmoney := money.(type) {
	// newmoney 在不同的 case 下回匹配到对应的类型
	case int:
		fmt.Println("money是 int", newmoney+11)
	case int64:
		fmt.Println("money是 int64", newmoney+22)
	case float64:
		fmt.Println("money是 float64", newmoney+33)
	case float32:
		fmt.Println("money是 float32", newmoney+44)
	default:
		fmt.Println("money是未知类型")
	}
	fmt.Println("结束", money)
}

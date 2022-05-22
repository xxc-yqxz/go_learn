package main

import "fmt"

func panicAndRecover() {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("系统出现严重故障", r)
	// 	}
	// }()
	// defer coverPanic() // 这样会报错，因为recover脱离了当前函数的调用栈
	// defer coverPanicUpdate() // 此种写法可行
	defer func() { // 此种写法也抓不住，因为recover 只会捕获当前函数上下文的错误，而此种写法相当于又嵌套了一层函数
		coverPanicUpdate()
	}()
	var nameScore map[string]int = nil

	nameScore["语文"] = 100 // 此处往空的map中赋值，会出现panic
}

func coverPanic() {
	func() {
		if r := recover(); r != nil {
			fmt.Println("系统出现严重故障", r)
		}
	}()
}

func coverPanicUpdate() {
	if r := recover(); r != nil {
		fmt.Println("系统出现严重故障", r)
	}
}

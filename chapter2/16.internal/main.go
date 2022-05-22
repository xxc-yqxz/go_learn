package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	startTime := time.Now()
	// defer 会在函数运行结束时被调用，并且可以有多个，根据定义的顺序从后往前运行
	// 并且即使程序有 panic，其中的defer也会被调用
	defer func() {
		finishTime := time.Now()
		fmt.Println(startTime, finishTime)
		fmt.Println(finishTime.Sub(startTime))
	}()
	defer func() {
		fmt.Println("123")
	}()
	arr := []int{}
	testPanic(arr)
	arr2 := make([]string, 3)
	arr3 := make([]string, 0, 4)                          // 定义 capacity 后，可以减少内存复制的性能损耗
	fmt.Println("len:%d", len(arr), "cap:%d", cap(arr))   // 0 0
	fmt.Println("len:%d", len(arr2), "cap:%d", cap(arr2)) // 3 3
	fmt.Println("len:%d", len(arr3), "cap:%d", cap(arr3)) // 0 4

	// 下面三行输出为空
	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])

	i := new(int)                  // new 会返回一个指针
	fmt.Println(reflect.TypeOf(i)) // *int

	arr4 := make([]string, 3, 4)
	arr5 := make([]string, 3, 4)
	// 原有复制切片写法
	// for i := 0; i < len(arr4); i++ {
	// 	arr5[i] = arr4[i]
	// }

	// 使用copy的写法
	copy(arr4, arr5)
}

func testPanic(arr []int) {
	// 下面三行报错
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
}

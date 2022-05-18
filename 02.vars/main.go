package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var hello string = "hello,golang!"
	var world = "world"
	fmt.Println(hello, world) // 此处虽然没有给 world 定义类型，但go会自动推测
	float1 := 1.24
	fmt.Println(float1)

	var int3, int4 uint = 33, 44 // 同时定义多个数据，并为每个数据设置类型（此时只需要在末尾设置一次）
	fmt.Println(int3, int4)

	//var ho, ver = 3, 4.12	// 报错
	var ho, ver float64 = 3, 4.12
	var sc = ho * ver // 初始化赋值时也可以直接传递表达式
	fmt.Println(sc)

	var heng, shu int
	fmt.Println(heng) // 0
	fmt.Println(shu)  // 0
	var name string
	fmt.Println(name) // 空
	var mapA map[string]string
	fmt.Println(mapA) // map[]

	a1 := "hello"
	fmt.Println(reflect.TypeOf(a1)) // string
	a2 := 3
	fmt.Println(reflect.TypeOf(a2)) // int
	a3 := 3.0
	fmt.Println(reflect.TypeOf(a3)) // float64
	a4 := &a3
	fmt.Println(reflect.TypeOf(a4)) // *float64

	var f1 float64 = 1.234
	var i1 int = int(f1)
	fmt.Println("f1:", f1, "i1:", i1) // f1: 1.234 i1: 1
	var a6 uint = math.MaxUint64
	var a7 int = int(a6)
	fmt.Println(a6, a7) // 18446744073709551615 -1
}

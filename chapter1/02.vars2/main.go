package main

import (
	"fmt"
	"strconv"
)

func main() {
	name := "123"
	name2 := "小强"
	{
		val, err := strconv.Atoi(name)
		fmt.Println(val, err) // 123 <nil>

		val2, err2 := strconv.Atoi(name2)
		fmt.Println(val2, err2) // 0 strconv.Atoi: parsing "小强": invalid syntax
	}
	age := 30
	fmt.Printf("%p\n", &age) // 0xc000018098
	age, time := 100, "123"  // 此处要求等号左边有新变量
	fmt.Printf("%p\n", &age) // 0xc000018098		--> 可以看到内存地址未变，只是修改了值
	fmt.Println(time)
	{
		age := 3
		fmt.Printf("%p\n", &age) // 0xc0000180d8
	}
}

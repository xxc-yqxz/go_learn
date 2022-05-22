package main

import "fmt"

func main() {
	{ //a := []int{}
		b := make([]int, 0) // 使用make 定义切片
		b = append(b, 1)
		fmt.Println("len:", len(b), b) // 1 [1]
	}
	{ //a := []int{}
		b := make([]int, 1) // 使用make 定义切片，第二个用来定义当前切片的初始长度，并给其赋予初始值，之后再添加会在其后面进行添加
		b = append(b, 1)
		fmt.Println("len:", len(b), b) // 2 [0 1]
	}
	{ //a := []int{}
		b := make([]int, 1, 2)                      // 第三个参数可以设置容量，这样会在开始时就创建空间，减少扩容时内存拷贝的性能损耗。
		fmt.Println("len:", len(b), "cap:", cap(b)) // len: 1 cap: 2
		b = append(b, 1)
		fmt.Println("len:", len(b), b) // 2 [0 1]
	}
}

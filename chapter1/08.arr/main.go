package main

import "fmt"

func main() {
	xqInfo := [3]string{"小强", "男", "在职"}
	xlInfo := [3]string{"小李", "男", "在职"}
	xsInfo := [3]string{"小苏", "女", "在职"}

	fmt.Println(xqInfo)
	fmt.Println(xlInfo)
	fmt.Println(xsInfo)

	// 难点：数组长度管理
	newPersonInfos := [3][3]string{
		[3]string{"小强", "男", "在职"},
		[3]string{"小李", "男", "在职"},
		[3]string{"小苏", "女", "在职"},
	}

	for _, val := range newPersonInfos {
		fmt.Println(val)
	}

	// ... 支持动态添加
	newPersonInfos2 := [...][3]string{
		[3]string{"小强", "男", "在职"},
		[3]string{"小李", "男", "在职"},
		[3]string{"小苏", "女", "在职"},
	}
	for _, val := range newPersonInfos2 {
		fmt.Println(val)
	}

	// 用降维方式输出
	newPersonInfos3 := [...][3]string{
		[3]string{"小强", "男", "在职"},
		[3]string{"小李", "男", "在职"},
		[3]string{"小苏", "女", "在职"},
	}
	for i, v1 := range newPersonInfos3 {
		for j, v2 := range v1 {
			fmt.Println(i, j, v2)
		}
	}
}

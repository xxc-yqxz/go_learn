package main

import "fmt"

// 此处定义的变量为全局变量，可以在任一作用域中使用
var allTall float64
var allWeight float64

func main() {
	tall, weight1 := 1.78, 78.0
	fmt.Println(tall, weight1)
	fmt.Println(allTall, allWeight)

	calculatorAdd := func(a, b float64) float64 {
		// 此处的 a、b 也只在函数体中可用
		return a + b
	}
	result := calculatorAdd(1, 3)
	fmt.Println(result)

	{
		personTall := 1.81
		personWeight := 90.0
		calculatorAdd(personTall, personWeight)
	}
	// fmt.Println(personTall)  // personTall 的有效范围为括号内部，外部无效

	sampleSubdomain()
	sampleSubdomain2()
}

func calcBMI() float64 {
	tall, weight2 := "1.78", 78.0 // 不同作用域可以定义同名变量
	// fmt.Println(weight1)		// 报错 -> 不可跨作用域访问变量
	fmt.Println(tall, weight2)
	return 0
}

func sampleSubdomain() {
	name := "小强"
	fmt.Println("名字是:", name)
	{
		fmt.Println(name) // 小强
		name = "Kr"
		fmt.Println(name) // Kr
	}
	fmt.Println(name) // Kr	上面使用赋值 "="，所以只是修改了上一作用域的name
}

func sampleSubdomain2() {
	name := "小强"
	fmt.Println("名字是:", name)
	{
		fmt.Println(name) // 小强
		name := "Kr"
		fmt.Println(name) // Kr
	}
	fmt.Println(name) // 小强	上面使用定义 ":="，所以在子作用域中新定义了个变量。不会影响父作用域的name

	if name == "小强" {
		// if else 中的 {} 也是一个单独的作用域
		a := 3
		fmt.Println(a)
	} else {
		a := 4
		fmt.Println(a)
	}
}

func sampleSubdomainIf() {
	if v := calcBMI(); v == 0 {
		fmt.Println(v)
	} else {
		fmt.Println(v)
	}
	// fmt.Println(v) // 无效，v的有效范围为
}

func sampleSubdomainFor() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello golang", i)
	}
	// fmt.Println(i)	// 无效
}

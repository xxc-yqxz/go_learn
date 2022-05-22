package main

import "fmt"

func main() {
	a := "您好"
	fmt.Println(a)
	aBytes := []byte(a)
	// ASCII 支持的字符编码有限，所以golang中使用了UTF8
	// 而byte方法只支持ASCII,会导致转换后的byte与原值的长度不同，此时可以使用rune方法
	fmt.Println(aBytes) // [230 130 168 229 165 189]
	bBytes := []rune(a)
	bBytes[0] = 'H'
	fmt.Println(bBytes, string(bBytes)) // [24744 22909]	 H好
	fmt.Println("修改切片内的内容")
	aBytes[0] = 'H' // 此处byte类型需要用 ''
	a = string(aBytes)
	fmt.Println(a) // H��好
}

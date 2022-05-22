package main

import "fmt"

func main() {
	hello()
	helloToSomeone("xxc")
	msg := constructHelloMessage("xxc")
	fmt.Println(msg)
}

func hello() {
	fmt.Println("你好，Golang")
}

func helloToSomeone(name string) {
	fmt.Println("你好,", name)
}

func constructHelloMessage(name string) string {
	return "你好，" + name
}

func calcBMI(tall float64, weight float64) float64 {
	return tall / (weight * weight)
}

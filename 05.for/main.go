package main

import "fmt"

func main() {
	fmt.Println("round 1")
	// i++ === i+=1 == i=i+1
	for i := 0; i < 5; i++ {
		fmt.Println("你好，golang")
	}

	fmt.Println("round 2")
	j := 0
	for ; j < 5; j++ {
		fmt.Println("round 2,hello,golang")
	}

	fmt.Println("round 3")
	k := 0
	for k < 5 {
		fmt.Println("round 3,hello,golang")
		k++
	}

	fmt.Println("round 4")
	i := 0
	for {
		fmt.Println("你好，Golang")
		i++
		if i > 3 {
			break
		}
	}

	fmt.Println("round 5")
	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println("你好，Golang")
	}
}

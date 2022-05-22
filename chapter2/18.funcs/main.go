package main

import "fmt"

func main() {
	panicAndRecover()
	closureMain()
	// openFile()
	fmt.Println(fib(10))
	fmt.Println(fib(9))
	// guess(1, 100)
	fmt.Println(calcSum(13, 3, 4, 23, 55, 66))
	fmt.Println(calcSum(13, 3, 4, 23, 55, 66))
	fmt.Println(calcSum(13, 3, 4, 23, 55, 66))
	fmt.Println(calcSum(13, 3, 4, 23, 55, 66))
	// showUsedTimes()
	// deferGuess()
}

func fib(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

package main

import (
	"fmt"
	"os"
	"time"
)

func deferGuess() {
	startTime := time.Now()
	fmt.Println("start time", startTime)
	// 此种写法不会得到5s后的结果
	defer fmt.Println("duration:", time.Now().Sub(startTime)) // defer后面的函数会先注册到内存栈中，而函数里面的内容都是已经计算好才放入的
	defer func() {
		fmt.Println("duration:", time.Now().Sub(startTime))
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("finish time", time.Now())
}

func openFile() {
	fileName := "/test.txt"
	fileObj, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	defer fileObj.Close()
}

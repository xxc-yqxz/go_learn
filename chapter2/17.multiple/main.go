package main

import "fmt"

func main() {
	//bmis := []float64{1.2, 3.222, 4.343443}
	avgBmi := calculateAvg(1, 2, 3, 4, 5, 6)
	fmt.Println(avgBmi)
}

func calculateAvg(bmis ...float64) (avgBmi float64) {
	fmt.Println(bmis) // [1 2 3 4 5 6]
	var total float64 = 0
	for _, item := range bmis {
		total += item
	}
	avgBmi = total / float64(len(bmis)) // 可以直接把 命名返回值进行赋值使用，并且不用在return处返回，它会自动返回
	return
}

func getScoresOfStudent(stdName string) (chinese int, match int, english int, physics int, nature int) {
	return 0, 0, 0, 0, 0
}

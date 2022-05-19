package main

import "fmt"

func main() {
	var weight float64 = 65.0
	var tall float64 = 1.78
	var bmi float64 = weight / (tall * tall)
	var age int = 35
	var sexWeight int
	var sex string = "男"
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	var fateRate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	fmt.Println("体脂率是:", fateRate)

	if sex == "男" {
		if age >= 18 && age <= 39 {
			if fateRate <= 0.1 {
				fmt.Println("目前偏瘦")
			} else if fateRate > 0.1 && fateRate <= 0.16 {
				fmt.Println("目前是：标准，要保持")
			} else if fateRate > 0.16 && fateRate <= 0.21 {
				fmt.Println("目前偏胖")
			} else if fateRate > 0.21 && fateRate <= 0.26 {
				fmt.Println("目前肥胖")
			} else {
				fmt.Println("过于肥胖")
			}
		} else if age >= 40 && age <= 59 {

		} else if age >= 60 {

		} else {
			fmt.Println("我们不看未成年人的体脂率，变化太大，无法评判")
		}
	} else {

	}
}

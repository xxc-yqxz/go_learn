package main

import "fmt"

func main() {
	for {
		var name string
		fmt.Print("姓名:")
		fmt.Scanln(&name)
		var weight float64
		fmt.Print("体重（kg）:")
		fmt.Scanln(&weight)
		var tall float64
		fmt.Print("身高（米）:")
		fmt.Scanln(&tall)
		var bmi float64 = weight / (tall * tall)
		var age int
		fmt.Print("年龄:")
		fmt.Scanln(&age)
		var sexWeight int
		var sex string = "男"
		fmt.Print("性别（男/女）:")
		fmt.Scanln(&sex)
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
			// todo:编写女性的体脂率与体脂状态表
		}

		var whetherContinue string
		fmt.Print("是否录入下一个(y/n)")
		fmt.Scanln(&whetherContinue)
		if whetherContinue != "y" {
			break
		}
	}
}

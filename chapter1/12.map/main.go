package main

import "fmt"

func main() {
	names := []string{"小强", "周毅", "训东"}
	fr := []float64{28, 18, 15}

	names = append(names, "xxc")
	fr = append(fr, 19)

	for i, name := range names {
		if name == "周毅" {
			fmt.Printf("%s 的体脂率是 %f\n", name, fr[i])
		}
	}

	var m1 map[string]int
	m2 := map[string]int{}
	m3 := map[string]int{"王强": 60, "李静": 83, "苗文": 91}
	fmt.Println(m1, m2, m3) // map[] map[] map[李静:83 王强:60 苗文:91]

	fmt.Println("王强的分数:", m3["王强"])
	xxcScore, ok := m3["xxc"] // 0 false
	fmt.Println(xxcScore, ok)
	m3["xxc"] = 123
	xxcScore, ok = m3["xxc"] // 123 true
	fmt.Println(xxcScore, ok)
	delete(m3, "xxc")
	m3["王强"] = 77
	fmt.Println(m3)

	for name, score := range m3 {
		fmt.Println(name, score)
	}

	var m4 map[string]int = nil
	// m4["a"] = 1 // 报错
	fmt.Println("m1没有实例化，可以直接取数字", m4["a"]) // 0
	delete(m4, "123")
}

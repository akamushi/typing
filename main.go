package main

import "fmt"

var example = []string{"golang", "hands-on", "in", "kagawa"}

func main() {
	var (
		data   []string
		ok, ng int
	)
	data = example

	for _, v := range data {
		fmt.Println(v)

		fmt.Print("input : ")
		var ans string
		fmt.Scanln(&ans)

		if v == ans {
			fmt.Println("○")
			ok++
		} else {
			fmt.Println("×")
			ng++
		}

	}

	fmt.Printf("[正誤] 正解: %d, 誤り: %d\n", ok, ng)
	cnt := float64(len(data))
	fmt.Printf("正解率: %f, 誤り率: %f\n", float64(ok)/cnt, float64(ng)/cnt)
}

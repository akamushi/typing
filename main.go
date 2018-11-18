package main

import (
	"fmt"
	"time"
)

var example = []string{"golang", "hands-on", "in", "kagawa"}

func main() {
	var (
		data   []string
		ok, ng int
	)
	data = example

	countdown()

	for _, v := range data {
		fmt.Println(v)

		fmt.Print("input : ")

		start := time.Now()
		var ans string
		fmt.Scanln(&ans)
		end := time.Now()

		fmt.Println(end.Sub(start))

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

func countdown() {
	for i := 0; i < 3; i++ {
		fmt.Println(3 - i)
		time.Sleep(time.Second * 1)
	}
}

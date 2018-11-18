package main

import "fmt"

var example = []string{"golang", "hands-on", "in", "kagawa"}

func main() {
	var data []string
	data = example

	for _, v := range data {
		fmt.Println(v)

		fmt.Print("input : ")
		var ans string
		fmt.Scanln(&ans)

		if v == ans {
			fmt.Println("○")
		} else {
			fmt.Println("×")
		}

	}
}

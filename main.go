package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var example = []string{"golang", "hands-on", "in", "kagawa"}

func main() {
	filename := "./example.txt"

	if _, err := os.Stat(filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	data, err := output(filename)
	if err != nil {
		os.Exit(1)
	}
	startGame(data)
}

func startGame(data []string) {
	var (
		ok, ng    int
		inputTime float64
	)
	countdown()

	for _, v := range data {
		fmt.Println(v)

		fmt.Print("input : ")

		start := time.Now()
		var ans string
		fmt.Scanln(&ans)
		end := time.Now()

		fmt.Println(end.Sub(start))
		inputTime += end.Sub(start).Seconds()

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
	fmt.Printf("平均入力時間: %f\n", float64(inputTime)/cnt)

}
func countdown() {
	for i := 0; i < 3; i++ {
		fmt.Println(3 - i)
		time.Sleep(time.Second * 1)
	}
}

func output(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var texts []string
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		text := scan.Text()
		if text != "" {
			texts = append(texts, text)
		}
	}
	return texts, nil
}

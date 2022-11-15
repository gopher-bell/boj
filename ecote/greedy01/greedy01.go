package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func getWord() int {
	scanner.Scan()
	r, _ := strconv.Atoi(scanner.Text())
	return r
}

func main() {
	money := getWord()
	coins := []int{500, 100, 50, 10}

	var count int
	for i := range coins {
		count += money / coins[i]
		money = money % coins[i]
	}

	writer.WriteString(strconv.Itoa(count))
	writer.Flush()
}

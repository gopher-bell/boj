package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func getWord() int {
	scanner.Scan()
	s, _ := strconv.Atoi(scanner.Text())
	return s
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	scanner.Split(bufio.ScanWords)

	n1 := getWord()

	arr := make([]int, n1)
	for i := range arr {
		arr[i] = getWord()
	}

	d := make([]int, n1)
	for i := range d {
		d[i] = 10000
	}

	d[0] = arr[0]
	d[1] = arr[1]

	for i := 2; i < len(arr); i++ {
		d[i] = max(d[i-1], d[i-2]+arr[i])
	}

	fmt.Println(d)
}

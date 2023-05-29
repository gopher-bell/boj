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

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanLines)

	res := 1
	for i := 0; i < 3; i++ {
		scanner.Scan()
		s := scanner.Text()
		r, _ := strconv.Atoi(s)
		res *= r
	}

	arr := make([]int, 10)
	for res > 10 {
		tmp := res % 10
		arr[tmp]++
		res /= 10
	}
	arr[res]++

	for i := range arr {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString("\n")
	}
}

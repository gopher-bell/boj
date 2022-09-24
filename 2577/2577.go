package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	var res = 1

	for i := 0; i < 3; i++ {
		scanner.Scan()
		s := scanner.Text()
		r, _ := strconv.Atoi(s)
		res *= r
	}

	s := strconv.Itoa(res)
	ss := strings.Split(s, "")

	ans := make([]int, 10)

	for _, v := range ss {
		k, _ := strconv.Atoi(v)
		ans[k]++
	}

	for i := range ans {
		writer.WriteString(strconv.Itoa(ans[i]))
		writer.WriteString("\n")
	}
}

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

func getWords() []int {
	scanner.Scan()
	s := strings.Split(scanner.Text(), "")
	res := make([]int, len(s))
	for i, v := range s {
		res[i], _ = strconv.Atoi(v)
	}
	return res
}

func main() {
	n1 := getWords()
	res := n1[0]

	for _, v := range n1 {
		if res <= 1 || v <= 1 {
			res += v
		} else {
			res *= v
		}
	}

	writer.WriteString(strconv.Itoa(res))
	writer.Flush()
}

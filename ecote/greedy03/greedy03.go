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
	if len(n1) == 1 {
		writer.WriteString(strconv.Itoa(n1[0]))
		writer.Flush()
		return
	}

	res := n1[0]

	for i := 1; i < len(n1); i++ {
		if res <= 1 || n1[i] <= 1 {
			res += n1[i]
		} else {
			res *= n1[i]
		}
	}

	writer.WriteString(strconv.Itoa(res))
	writer.Flush()
}

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

func getWord() int {
	scanner.Scan()
	s, _ := strconv.Atoi(scanner.Text())
	return s
}

func isThree(i, j, k int) bool {
	s1 := strconv.Itoa(i)
	s2 := strconv.Itoa(j)
	s3 := strconv.Itoa(k)

	return strings.Contains(s1+s2+s3, "3")
}

func main() {
	n := getWord()
	var cnt int
	for i := 0; i <= n; i++ {
		for j := 0; j <= 59; j++ {
			for k := 0; k <= 59; k++ {
				if isThree(i, j, k) {
					cnt++
				}
			}
		}
	}

	writer.WriteString(strconv.Itoa(cnt))
	writer.Flush()
}

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
	dx      = []int{-2, -1, 1, 2, 2, 1, -1, -2}
	dy      = []int{-1, -2, -2, -1, 1, 2, 2, 1}
)

func getWords() (int, int) {
	scanner.Scan()
	s := strings.Split(scanner.Text(), "")
	s1 := []rune(s[0])
	s2 := s1[0] - 'a' + 1
	s3, _ := strconv.Atoi(s[1])

	return int(s2), s3
}

func main() {
	n1, n2 := getWords()
	var res int
	for i := 0; i < 8; i++ {
		nx := n2 - dx[i]
		ny := n1 - dy[i]
		if nx >= 1 && nx <= 7 && ny >= 1 && ny <= 8 {
			res++
		}
	}
	writer.WriteString(strconv.Itoa(res))
	writer.Flush()
}

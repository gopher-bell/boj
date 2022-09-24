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
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	texts := make([]string, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		texts[i] = s
	}

	for i := range texts {
		var score = 1
		var total int
		for j := range texts[i] {
			if texts[i][j] == 'O' {
				total += score
				score++
			} else {
				score = 1
			}
		}
		writer.WriteString(strconv.Itoa(total))
		writer.WriteString("\n")
	}
}

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
	n1, n2 := getWord(), getWord()
	var count int

	for {
		if n1 == 1 {
			break
		} else if n1%n2 == 0 {
			n1 /= n2
			count++
		} else {
			n1--
			count++
		}
	}

	writer.WriteString(strconv.Itoa(count))
	writer.Flush()
}

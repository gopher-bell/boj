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
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	s, _ := strconv.Atoi(scanner.Text())

	var n = 1
	var p = 1
	for {
		if n >= s {
			break
		}
		n += 6 * p
		p++
	}
	writer.WriteString(strconv.Itoa(p))
}

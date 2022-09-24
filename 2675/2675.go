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
	scanner.Split(bufio.ScanLines)

	var n int
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		ss := strings.Fields(s)
		num, _ := strconv.Atoi(ss[0])
		for j := range ss[1] {
			for k := 0; k < num; k++ {
				writer.WriteString(string(ss[1][j]))
			}
		}
		writer.WriteString("\n")
	}
}

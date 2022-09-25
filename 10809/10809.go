package main

import (
	"bufio"
	"bytes"
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
	b := bytes.Fields(scanner.Bytes())
	comp := make([]int, 26)
	for i := range comp {
		comp[i] = -1
	}
	res := make([]byte, len(b[0]))
	for i := range b[0] {
		res[i] = b[0][i]
	}
	for i := range res {
		if comp[res[i]-97] == -1 {
			comp[res[i]-97] = i
		}
	}
	for i := range comp {
		writer.WriteString(strconv.Itoa(comp[i]) + " ")
	}
}

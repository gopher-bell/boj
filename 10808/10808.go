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
	t := scanner.Text()
	arr := make([]int, 26)

	for i := range t {
		arr[t[i]-'a']++
	}

	for i, v := range arr {
		writer.WriteString(strconv.Itoa(v))
		if i != len(arr)-1 {
			writer.WriteString(" ")
		}
	}
}

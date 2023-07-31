package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	n, m    int
	arr     [10]int
	visited [10]bool
)

func getNumber() int {
	ret := 0
	scanner.Scan()
	for _, v := range scanner.Bytes() {
		ret *= 10
		ret += int(v - '0')
	}

	return ret
}

func fn(k int) {
	if k == m {
		for i := 0; i < m; i++ {
			writer.WriteString(strconv.Itoa(arr[i]) + " ")
		}
		writer.WriteString("\n")

		return
	}

	for i := 1; i <= n; i++ {
		if !visited[i] {
			visited[i] = true
			arr[k] = i
			fn(k + 1)
			visited[i] = false
		}
	}
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = getNumber(), getNumber()
	fn(0)
}

package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	graph    [][]int
	visited1 []bool
	visited2 []bool
	visited3 []bool
	n        int
	cnt      int
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

func solution(current int) {
	if current == n {
		cnt++
		return
	}

	for i := 0; i < n; i++ {
		if visited1[i] || visited2[i+current] || visited3[current-i+n-1] {
			continue
		}
		visited1[i] = true
		visited2[i+current] = true
		visited3[current-i+n-1] = true
		solution(current + 1)
		visited1[i] = false
		visited2[i+current] = false
		visited3[current-i+n-1] = false
	}
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n = getNumber()
	graph = make([][]int, n)
	visited1 = make([]bool, n*2)
	visited2 = make([]bool, n*2)
	visited3 = make([]bool, n*2)
	for i := range graph {
		graph[i] = make([]int, n)
	}
	solution(0)
	writer.WriteString(strconv.Itoa(cnt))
}

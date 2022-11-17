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
	dx      = []int{1, -1, 0, 0}
	dy      = []int{0, 0, -1, 1}
	result  = 0
	n, m    int
)

func getWords() []int {
	if scanner.Scan() {
		res := make([]int, m)
		s := strings.Split(scanner.Text(), "")
		for i := range s {
			res[i], _ = strconv.Atoi(s[i])
		}
		return res
	}

	return nil
}

func dfs(graph [][]int, y, x int) bool {
	if graph[y][x] == 0 {
		graph[y][x] = 1
		for k := 0; k < 4; k++ {
			ny := y + dy[k]
			nx := x + dx[k]
			if ny > 0 && ny < n && nx > 0 && nx < m {
				dfs(graph, ny, nx)
			}
		}
		return true
	}

	return false
}

func main() {
	scanner.Scan()
	s := strings.Fields(scanner.Text())
	n, _ = strconv.Atoi(s[0])
	m, _ = strconv.Atoi(s[1])

	graph := make([][]int, n)

	for i := range graph {
		graph[i] = getWords()
	}

	for i := range graph {
		for j := range graph[i] {
			if dfs(graph, i, j) {
				result++
			}
		}
	}

	writer.WriteString(strconv.Itoa(result))
	writer.Flush()
}

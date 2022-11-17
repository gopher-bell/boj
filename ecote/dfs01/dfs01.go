package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	dx      = []int{1, -1, 0, 0}
	dy      = []int{0, 0, -1, 1}
	result  = 0
	n, m    int
)

func getWords() int {
	if scanner.Scan() {
		s, _ := strconv.Atoi(scanner.Text())
		return s
	}

	return -1
}

func dfs(graph [][]int, y, x int) bool {
	if graph[y][x] == 0 {
		graph[y][x] = 1
		for k := 0; k < 4; k++ {
			ny := y + dy[y]
			nx := x + dx[x]
			if ny > 0 && ny < n && nx > 0 && nx < m {
				dfs(graph, ny, nx)
			}
		}
	}

	return false
}

func main() {
	scanner.Split(bufio.ScanBytes)

	n, m = getWords(), getWords()

	graph := make([][]int, n)
	visited := make([][]bool, n)

	for i := range graph {
		graph[i] = make([]int, m)
		visited[i] = make([]bool, m)
	}

	for i := range graph {
		for j := range graph[i] {
			graph[i][j] = getWords()
		}
	}

	for i := range graph {
		for j := range graph[i] {
			if dfs(graph, i, j) {
				result++
			}
		}
	}
}

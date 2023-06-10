package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	scanner    = bufio.NewScanner(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
	dx         = []int{1, 0, -1, 0}
	dy         = []int{0, 1, 0, -1}
	queueSlice []queue
	n, m       int
)

type queue struct {
	x int
	y int
}

func getNumber() int {
	ret := 0
	flag := false
	scanner.Scan()
	for _, v := range scanner.Bytes() {
		if v == '-' {
			flag = true
			continue
		}
		ret *= 10
		ret += int(v - '0')
	}

	if flag {
		return -ret
	}

	return ret
}

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func bfs(graph [][]int) {
	for len(queueSlice) > 0 {
		dq := queueSlice[0]
		queueSlice = queueSlice[1:]
		for i := 0; i < 4; i++ {
			nx := dq.x + dx[i]
			ny := dq.y + dy[i]
			if nx < 0 || nx >= n || ny < 0 || ny >= m {
				continue
			}
			if graph[nx][ny] != 0 {
				continue
			}
			graph[nx][ny] = graph[dq.x][dq.y] + 1
			queueSlice = append(queueSlice, queue{nx, ny})

		}
	}
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	m, n = getNumber(), getNumber()

	graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, m)
		for j := range graph[i] {
			graph[i][j] = getNumber()
			if graph[i][j] == 1 {
				queueSlice = append(queueSlice, queue{i, j})
			}
		}
	}

	bfs(graph)

	max := 0
	for i := range graph {
		for j := range graph[i] {
			if graph[i][j] == 0 {
				writer.WriteString("-1")
				return
			}
			max = getMax(max, graph[i][j])
		}
	}

	writer.WriteString(strconv.Itoa(max - 1))
}

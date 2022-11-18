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
	dy      = []int{0, 0, 1, -1}
	dx      = []int{1, -1, 0, 0}
	n1      int
	n2      int
	level   = 1
)

type node struct {
	y int
	x int
}

func bfs(graph [][]int, y, x int) {
	queue := []node{{y, x}}
	for len(queue) != 0 {
		elem := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			ny := elem.y + dy[i]
			nx := elem.x + dx[i]
			if 0 <= ny && ny < n1 && 0 <= nx && nx < n2 {
				if graph[ny][nx] == 1 {
					graph[ny][nx] = graph[elem.y][elem.x] + 1
					queue = append(queue, node{ny, nx})
				}
			}
		}
	}
}

func main() {
	scanner.Scan()
	s := strings.Fields(scanner.Text())
	n1, _ = strconv.Atoi(s[0])
	n2, _ = strconv.Atoi(s[1])

	graph := make([][]int, n1)
	visited := make([][]int, n1)
	for i := range graph {
		graph[i] = make([]int, n2)
		visited[i] = make([]int, n2)
		scanner.Scan()
		t := strings.Split(scanner.Text(), "")
		for j := range graph[i] {
			graph[i][j], _ = strconv.Atoi(t[j])
		}
	}

	bfs(graph, 0, 0)
	writer.WriteString(strconv.Itoa(graph[n1-1][n2-1]))
	writer.Flush()
}

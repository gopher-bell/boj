package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	dx      = []int{1, 0, -1, 0}
	dy      = []int{0, 1, 0, -1}
)

type pair struct {
	x int
	y int
}

func getNumber() int {
	scanner.Scan()
	r := 0
	for _, c := range scanner.Bytes() {
		r *= 10
		r += int(c - '0')
	}
	return r
}

func bfs(graph [][]int, visited [][]bool, x, y int) (int, bool) {
	ret := false
	cnt := 1
	queue := []pair{{x, y}}
	for len(queue) > 0 {
		dq := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			nx := dx[i] + dq.x
			ny := dy[i] + dq.y
			if nx < len(graph) && nx >= 0 && ny < len(graph[0]) && ny >= 0 {
				if !visited[nx][ny] && graph[nx][ny] == 1 {
					visited[nx][ny] = true
					queue = append(queue, pair{nx, ny})
					cnt++
				}
			}
		}
		ret = true
	}

	return cnt, ret
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n1, n2 := getNumber(), getNumber()
	graph := make([][]int, n1)
	visited := make([][]bool, n1)
	for i := range graph {
		visited[i] = make([]bool, n2)
		graph[i] = make([]int, n2)
		for j := range graph[i] {
			graph[i][j] = getNumber()
		}
	}

	cnt := 0
	max := 0
	for i := range visited {
		for j := range visited[i] {
			if !visited[i][j] && graph[i][j] == 1 {
				visited[i][j] = true
				tot, success := bfs(graph, visited, i, j)
				if success {
					cnt++
				}
				if tot > max {
					max = tot
				}
			}
		}
	}

	writer.WriteString(strconv.Itoa(cnt) + "\n" + strconv.Itoa(max))
}

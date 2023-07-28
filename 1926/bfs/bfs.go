package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	writer  = bufio.NewWriter(os.Stdout)
	scanner = bufio.NewScanner(os.Stdin)
	dx      = []int{1, 0, -1, 0}
	dy      = []int{0, 1, 0, -1}
)

type q struct {
	x int
	y int
}

func transferSlice(s []string, rg int) []int {
	n := make([]int, rg)
	for i := 0; i < rg; i++ {
		n[i], _ = strconv.Atoi(s[i])
	}

	return n
}

func bfs(graph [][]int, x, y int) int {
	tot := 1
	queue := []q{{x, y}}
	for len(queue) > 0 {
		dq := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			nx, ny := dq.x+dx[i], dq.y+dy[i]
			if nx < 0 || nx >= len(graph) || ny < 0 || ny >= len(graph[0]) {
				continue
			}
			if graph[nx][ny] != 1 {
				continue
			}

			graph[nx][ny] = 0
			queue = append(queue, q{nx, ny})
			tot++
		}
	}

	return tot
}

func main() {
	defer writer.Flush()
	scanner.Scan()
	s := strings.Split(scanner.Text(), " ")
	n1, _ := strconv.Atoi(s[0])
	n2, _ := strconv.Atoi(s[1])

	graph := make([][]int, n1)
	for i := range graph {
		scanner.Scan()
		s2 := strings.Split(scanner.Text(), " ")
		graph[i] = transferSlice(s2, n2)
	}

	cnt := 0
	max := 0

	for i := range graph {
		for j := range graph[i] {
			if graph[i][j] == 1 {
				graph[i][j] = 0
				cnt++
				res := bfs(graph, i, j)
				if res > max {
					max = res
				}
			}
		}
	}

	writer.WriteString(strconv.Itoa(cnt))
	writer.WriteString("\n")
	writer.WriteString(strconv.Itoa(max))
}

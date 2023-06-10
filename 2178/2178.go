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
	dx      = []int{1, 0, -1, 0}
	dy      = []int{0, 1, 0, -1}
)

type queue struct {
	x int
	y int
}

func getNumbers() []int {
	scanner.Scan()
	nums := make([]int, 0)
	for _, v := range scanner.Bytes() {
		nums = append(nums, int(v-'0'))
	}

	return nums
}

func bfs(graph [][]int, x, y int) {
	q := []queue{{x: x, y: y}}
	for len(q) > 0 {
		dq := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			nx := dq.x + dx[i]
			ny := dq.y + dy[i]
			if nx < 0 || nx >= len(graph) || ny < 0 || ny >= len(graph[0]) || graph[nx][ny] != 1 {
				continue
			}
			graph[nx][ny] = graph[dq.x][dq.y] + 1
			q = append(q, queue{x: nx, y: ny})
		}
	}
}

func main() {
	defer writer.Flush()
	scanner.Scan()
	input := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(input[0])

	graph := make([][]int, n)
	for i := range graph {
		scanner.Scan()
		s := scanner.Text()
		res := make([]int, len(s))
		for j := range s {
			k, _ := strconv.Atoi(string(s[j]))
			res[j] = k
		}
		graph[i] = res
	}

	bfs(graph, 0, 0)

	writer.WriteString(strconv.Itoa(graph[len(graph)-1][len(graph[0])-1]))
}

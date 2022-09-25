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
	graph   [][]int
	weight  [][]int
	n       int
	m       int
)

func bfs(y, x int) {
	queue := [][]int{{y, x}}
	weight[y][x] = 1
	for len(queue) != 0 {
		dq := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			ey := dq[0] + dy[i]
			ex := dq[1] + dx[i]

			if (0 <= ey && ey < n) && (0 <= ex && ex < m) {
				if graph[ey][ex] == 1 && weight[ey][ex] == 0 {
					weight[ey][ex] = weight[dq[0]][dq[1]] + 1
					queue = append(queue, []int{ey, ex})
				}
			}
		}
	}
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	input := strings.Fields(scanner.Text())
	n, _ = strconv.Atoi(input[0])
	m, _ = strconv.Atoi(input[1])

	graph = make([][]int, n)
	weight = make([][]int, n)

	for i := range graph {
		scanner.Scan()
		s := scanner.Text()
		s = s[:m]
		res := make([]int, len(s))
		for j := range s {
			k, _ := strconv.Atoi(string(s[j]))
			res[j] = k
		}
		graph[i] = res
		weight[i] = make([]int, m)
	}

	bfs(0, 0)
	writer.WriteString(strconv.Itoa(weight[n-1][m-1]))
}

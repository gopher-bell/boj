package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	dx        = []int{1, 0, -1, 0}
	dy        = []int{0, 1, 0, -1}
	graph     [][]int
	check     [][]bool
	totalDraw int
	MaxDraw   int
)

func readStdin() []int {
	scanner.Scan()
	s := strings.Fields(scanner.Text())
	res := make([]int, len(s))
	for i := range s {
		t, _ := strconv.Atoi(s[i])
		res[i] = t
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func bfs(y, x int) int {
	var res = 1
	queue := [][]int{{y, x}}
	for len(queue) != 0 {
		dq := queue[0]
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			ny := dq[0] + dy[i]
			nx := dq[1] + dx[i]
			if (0 <= ny && ny < len(graph)) && (0 <= nx && nx < len(graph[0])) {
				if graph[ny][nx] == 1 && !check[ny][nx] {
					check[ny][nx] = true
					res++
					queue = append(queue, []int{ny, nx})
				}
			}
		}
	}

	return res
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanLines)

	n1 := readStdin()

	graph = make([][]int, n1[0])
	check = make([][]bool, n1[0])

	for i := 0; i < n1[0]; i++ {
		graph[i] = readStdin()
		check[i] = make([]bool, n1[1])
	}

	for i := range graph {
		for j := range graph[i] {
			if graph[i][j] == 1 && !check[i][j] {
				check[i][j] = true
				totalDraw++
				MaxDraw = max(MaxDraw, bfs(i, j))
			}
		}
	}

	writer.WriteString(strconv.Itoa(totalDraw))
	writer.WriteString("\n")
	writer.WriteString(strconv.Itoa(MaxDraw))
}

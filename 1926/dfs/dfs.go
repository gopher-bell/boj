package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	graph   [][]int
	visited [][]bool
	n       int
	m       int
	area    int
	maxArea int
	dy      = []int{1, 0, -1, 0}
	dx      = []int{0, 1, 0, -1}
)

func readStdin() (int, error) {
	scanner.Scan()
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, err
	}

	return num, nil
}

func dfs(y, x int) {
	area++
	if area > maxArea {
		maxArea = area
	}
	for i := 0; i < 4; i++ {
		ny := y + dy[i]
		nx := x + dx[i]

		if (0 <= ny && ny < n) && (0 <= nx && nx < m) {
			if graph[ny][nx] == 1 && !visited[ny][nx] {
				visited[ny][nx] = true
				dfs(ny, nx)

			}
		}
	}

}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	var err error
	n, err = readStdin()
	if err != nil {
		panic(err)
	}
	m, err = readStdin()
	if err != nil {
		panic(err)
	}

	graph = make([][]int, n)
	visited = make([][]bool, n)

	for i := range graph {
		graph[i] = make([]int, m)
		visited[i] = make([]bool, m)
		for j := range graph[i] {
			graph[i][j], err = readStdin()
			if err != nil {
				panic(err)
			}
		}
	}

	var total int

	for i := range graph {
		for j := range graph[i] {
			if graph[i][j] == 1 && !visited[i][j] {
				visited[i][j] = true
				total++
				area = 0
				dfs(i, j)
			}
		}
	}

	writer.WriteString(strconv.Itoa(total))
	writer.WriteString("\n")
	writer.WriteString(strconv.Itoa(maxArea))
}

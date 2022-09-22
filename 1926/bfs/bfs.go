package main

import (
	"bufio"
	"fmt"
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

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func bfs(y, x int) int {
	queue := [][]int{{y, x}}
	var count = 1
	for len(queue) != 0 {
		dq := queue[0]
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			ny := dq[0] + dy[i]
			nx := dq[1] + dx[i]

			if (0 <= ny && ny < n) && (0 <= nx && nx < m) {
				if graph[ny][nx] == 1 && !visited[ny][nx] {
					visited[ny][nx] = true
					count++
					queue = append(queue, []int{ny, nx})
				}
			}
		}
	}

	return count
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

	var total, biggest int

	for i := range graph {
		for j := range graph[i] {
			if graph[i][j] == 1 && !visited[i][j] {
				visited[i][j] = true
				total++
				biggest = max(biggest, bfs(i, j))
			}
		}
	}

	fmt.Println(total)
	fmt.Println(biggest)
}

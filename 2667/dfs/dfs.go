package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	dy      = []int{1, 0, -1, 0}
	dx      = []int{0, 1, 0, -1}
	graph   [][]int
	visited [][]bool
	n       int
	res     []int
	max     int
)

func readStdin() ([]int, error) {
	scanner.Scan()
	var nums []int
	s := scanner.Text()
	for i := range s {
		num, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}

	return nums, nil
}

func readStdinOne() (int, error) {
	scanner.Scan()
	s := scanner.Text()

	num, err := strconv.Atoi(s)
	if err != nil {
		return num, err
	}

	return num, nil
}

func dfs(y, x int) {
	max++
	for i := 0; i < 4; i++ {
		ny := y + dy[i]
		nx := x + dx[i]

		if (0 <= ny && ny < n) && (0 <= nx && nx < n) {
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
	n, err = readStdinOne()
	if err != nil {
		panic(err)
	}

	graph = make([][]int, n)
	visited = make([][]bool, n)

	for i := range graph {
		graph[i] = make([]int, n)
		visited[i] = make([]bool, n)
		graph[i], err = readStdin()
		if err != nil {
			panic(err)
		}
	}

	for i := range graph {
		for j := range graph[i] {
			if graph[i][j] == 1 && !visited[i][j] {
				visited[i][j] = true
				max = 0
				dfs(i, j)
				res = append(res, max)
			}
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	writer.WriteString(strconv.Itoa(len(res)))
	writer.WriteString("\n")
	for i := range res {
		writer.WriteString(strconv.Itoa(res[i]))
		writer.WriteString("\n")
	}
}

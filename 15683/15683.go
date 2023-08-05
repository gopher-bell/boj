package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	writer  = bufio.NewWriter(os.Stdout)
	scanner = bufio.NewScanner(os.Stdin)
	graph   [][]int
)

func max(nums ...int) int {
	num := 0
	for _, v := range nums {
		if v > num {
			num = v
		}
	}

	return num
}

func printLeft(x, y int) {
	for j := 0; j < y; j++ {
		if graph[x][j] == 0 {
			graph[x][j] = -1
		}
	}
}

func printRight(x, y int) {
	for j := y + 1; j < len(graph[0]); j++ {
		if graph[x][j] == 0 {
			graph[x][j] = -1
		}
	}
}

func printTop(x, y int) {
	for i := 0; i < x; i++ {
		if graph[i][y] == 0 {
			graph[i][y] = -1
		}
	}
}

func printBottom(x, y int) {
	for i := x + 1; i < len(graph); i++ {
		if graph[i][y] == 0 {
			graph[i][y] = -1
		}
	}
}

func searchLine(x, y int) (int, int, int, int) {
	left := 0
	right := 0
	top := 0
	bottom := 0

	for j := 0; j < y; j++ {
		if graph[x][j] == 6 {
			break
		}
		if graph[x][j] == 0 {
			left++
		}
	}

	for j := y + 1; j < len(graph[0]); j++ {
		if graph[x][j] == 6 {
			break
		}
		if graph[x][j] == 0 {
			right++
		}
	}

	for i := 0; i < x; i++ {
		if graph[i][y] == 6 {
			break
		}
		if graph[i][y] == 0 {
			top++
		}
	}

	for i := x + 1; i < len(graph); i++ {
		if graph[i][y] == 6 {
			break
		}
		if graph[i][y] == 0 {
			bottom++
		}
	}

	return left, right, top, bottom
}

func solution() {
	for i := range graph {
		for j := range graph {
			if graph[i][j] > 0 {
				// left, right, top, bottom := searchLine(i, j)
				switch graph[i][j] {
				case 1:
				case 2:
				case 3:
				case 4:
				case 5:
				}
			}
		}
	}
}

func main() {
	graph = [][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 6, 0},
		{0, 0, 0, 0, 0, 0},
	}

	l, r, t, b := searchLine(2, 2)

	fmt.Println(l, r, t, b)
}

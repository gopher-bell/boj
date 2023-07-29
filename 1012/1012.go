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

type queue struct {
	x, y int
}

func getNumber() int {
	ret := 0
	scanner.Scan()
	for _, v := range scanner.Bytes() {
		ret *= 10
		ret += int(v - '0')
	}

	return ret
}

func bfs(farm [][]int, x, y int) {
	q := []queue{{x, y}}
	for len(q) > 0 {
		dq := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			nx, ny := dq.x+dx[i], dq.y+dy[i]
			if nx < 0 || nx >= len(farm) || ny < 0 || ny >= len(farm[0]) {
				continue
			}
			if farm[nx][ny] != 1 {
				continue
			}

			farm[nx][ny]++
			q = append(q, queue{nx, ny})
		}
	}
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	tests := getNumber()

	for i := 0; i < tests; i++ {
		cnt := 0
		cols, rows, nums := getNumber(), getNumber(), getNumber()
		farm := make([][]int, rows)
		for j := range farm {
			farm[j] = make([]int, cols)
		}
		for k := 0; k < nums; k++ {
			y, x := getNumber(), getNumber()
			farm[x][y] = 1
		}
		for m := range farm {
			for n := range farm[m] {
				if farm[m][n] == 1 {
					cnt++
					farm[m][n]++
					bfs(farm, m, n)
				}
			}
		}

		writer.WriteString(strconv.Itoa(cnt))
		if i != tests-1 {
			writer.WriteString("\n")
		}
	}
}

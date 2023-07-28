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

func bfs(BGraph [][]string, FGraph, JGraph [][]int) int {
	FQueue := make([]q, 0)
	JQueue := make([]q, 0)
	for i := range BGraph {
		for j := range BGraph[i] {
			if BGraph[i][j] == "F" {
				FQueue = append(FQueue, q{i, j})
				FGraph[i][j] = 0
			} else if BGraph[i][j] == "J" {
				JQueue = append(JQueue, q{i, j})
				JGraph[i][j] = 0
			}
		}
	}

	for len(FQueue) > 0 {
		dq := FQueue[0]
		FQueue = FQueue[1:]
		for i := 0; i < 4; i++ {
			nx, ny := dq.x+dx[i], dq.y+dy[i]
			if nx < 0 || nx >= len(FGraph) || ny < 0 || ny >= len(FGraph[0]) {
				continue
			}
			if FGraph[nx][ny] >= 0 || BGraph[nx][ny] == "#" {
				continue
			}

			FGraph[nx][ny] = FGraph[dq.x][dq.y] + 1
			FQueue = append(FQueue, q{nx, ny})
		}
	}

	for len(JQueue) > 0 {
		dq := JQueue[0]
		JQueue = JQueue[1:]
		if dq.x == 0 || dq.y == 0 || dq.x == len(JGraph)-1 || dq.y == len(JGraph[0])-1 {
			return JGraph[dq.x][dq.y] + 1
		}
		for i := 0; i < 4; i++ {
			nx, ny := dq.x+dx[i], dq.y+dy[i]
			if nx < 0 || nx >= len(JGraph) || ny < 0 || ny >= len(JGraph[0]) {
				continue
			}
			if JGraph[nx][ny] >= 0 || BGraph[nx][ny] == "#" {
				continue
			}
			if FGraph[nx][ny] != -1 && FGraph[nx][ny] <= JGraph[dq.x][dq.y]+1 {
				continue
			}

			JGraph[nx][ny] = JGraph[dq.x][dq.y] + 1
			JQueue = append(JQueue, q{nx, ny})
		}
	}

	return -1
}

func main() {
	defer writer.Flush()
	scanner.Scan()
	s := strings.Split(scanner.Text(), " ")
	n1, _ := strconv.Atoi(s[0])
	n2, _ := strconv.Atoi(s[1])

	BGraph := make([][]string, n1)
	FGraph := make([][]int, n1)
	JGraph := make([][]int, n1)
	for i := range BGraph {
		scanner.Scan()
		s2 := strings.Split(scanner.Text(), "")
		s2 = s2[:n2]
		BGraph[i] = s2

		FGraph[i] = make([]int, n2)
		JGraph[i] = make([]int, n2)
		for j := 0; j < n2; j++ {
			FGraph[i][j] = -1
			JGraph[i][j] = -1
		}
	}

	res := bfs(BGraph, FGraph, JGraph)
	if res == -1 {
		writer.WriteString("IMPOSSIBLE")
		return
	}

	writer.WriteString(strconv.Itoa(res))
}

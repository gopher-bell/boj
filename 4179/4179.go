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
	dx      = [4]int{1, 0, -1, 0}
	dy      = [4]int{0, 1, 0, -1}
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

func getString() []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), "")
}

func bfs(graph [][]string, queueSlice []queue) (bool, int) {
	cnt := 1
	success := false
	for len(queueSlice) > 0 {
		dq := queueSlice[0]
		queueSlice = queueSlice[1:]
		for i := 0; i < 4; i++ {
			nx := dq.x + dx[i]
			ny := dq.y + dy[i]
			// 범위에 안맞으면 진행 x
			if nx < 0 || nx > len(graph)-1 || ny < 0 || ny > len(graph[0])-1 {
				continue
			}
			// 길이 아니면 진행 x
			if graph[nx][ny] != "." {
				continue

			}

			if graph[dq.x][dq.y] == "J" {
				// 불과 겹치면 진행 x
				for j := 0; j < 4; j++ {
					nnx := nx + dx[j]
					nny := ny + dy[j]
					if nnx < 0 || nnx > len(graph)-1 || nny < 0 || nny > len(graph[0])-1 {
						continue
					}
					if graph[nnx][nny] == "F" {
						continue
					}
				}

				cnt++

				if nx == 0 || nx == len(graph)-1 || ny == 0 || ny == len(graph[0])-1 {
					success = true
					break
				}
			}

			graph[nx][ny] = graph[dq.x][dq.y]
			queueSlice = append(queueSlice, queue{nx, ny})
		}
	}

	return success, cnt
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m := getNumber(), getNumber()
	queueSlice := make([]queue, 0, n*m)
	graph := make([][]string, n)
	for i := range graph {
		graph[i] = make([]string, m)
		graph[i] = getString()
		for j := range graph[i] {
			if graph[i][j] == "J" || graph[i][j] == "F" {
				queueSlice = append(queueSlice, queue{i, j})
			}
		}
	}

	success, cnt := bfs(graph, queueSlice)
	if success {
		writer.WriteString(strconv.Itoa(cnt))
		return
	}

	writer.WriteString("IMPOSSIBLE")
}

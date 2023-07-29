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

func bfs(graph [][]string, visited [][]bool, x, y int, rule string) {
	q := []queue{{x, y}}
	for len(q) > 0 {
		dq := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			nx, ny := dq.x+dx[i], dq.y+dy[i]
			if nx < 0 || nx >= len(graph) || ny < 0 || ny >= len(graph[0]) {
				continue
			}
			if visited[nx][ny] {
				continue
			}
			if graph[nx][ny] != rule {
				continue
			}

			visited[nx][ny] = true
			q = append(q, queue{nx, ny})
		}
	}
}

func bfsGB(graph [][]string, visited [][]bool, x, y int, color string) {
	q := []queue{{x, y}}
	for len(q) > 0 {
		dq := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			nx, ny := dq.x+dx[i], dq.y+dy[i]
			if nx < 0 || nx >= len(graph) || ny < 0 || ny >= len(graph[0]) {
				continue
			}
			if visited[nx][ny] {
				continue
			}
			if color == "B" {
				if graph[nx][ny] != color {
					continue
				}
			} else {
				if graph[nx][ny] != "R" && graph[nx][ny] != "G" {
					continue
				}
			}

			visited[nx][ny] = true
			q = append(q, queue{nx, ny})
		}
	}
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := getNumber()

	graph := make([][]string, n)
	visited := make([][]bool, n)
	visitedRGB := make([][]bool, n)
	for i := range graph {
		scanner.Scan()
		s := strings.Split(scanner.Text(), "")
		graph[i] = s
		visited[i] = make([]bool, n)
		visitedRGB[i] = make([]bool, n)
	}

	rule := [3]string{"R", "G", "B"}
	cnt := [3]int{0, 0, 0}
	cnt2 := 0
	for i := range rule {
		for j := range graph {
			for k := range graph[j] {
				if !visited[j][k] && graph[j][k] == rule[i] {
					cnt[i]++
					visited[j][k] = true
					bfs(graph, visited, j, k, rule[i])
				}
			}

			for k := range graph[i] {
				if !visitedRGB[j][k] {
					if !visitedRGB[j][k] && graph[j][k] == rule[i] {
						cnt2++
						visitedRGB[j][k] = true
						bfsGB(graph, visitedRGB, j, k, rule[i])
					}
				}
			}
		}
	}

	writer.WriteString(strconv.Itoa(cnt[0]+cnt[1]+cnt[2]) + " " + strconv.Itoa(cnt2))
}

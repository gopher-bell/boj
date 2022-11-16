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
	dy        = []int{0, 0, 1, -1}
	dx        = []int{1, -1, 0, 0}
	moveTypes = []string{"R", "L", "D", "U"}
)

func getWord() int {
	scanner.Scan()
	r, _ := strconv.Atoi(scanner.Text())
	return r
}

func getWords() []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), " ")
}

func main() {
	n1 := getWord()
	cmd := getWords()

	graph := make([][]int, n1)

	for i := range graph {
		graph[i] = make([]int, n1)
	}

	x, y := 1, 1
	for i := range cmd {
		var nx, ny int
		for j := range moveTypes {
			if cmd[i] == moveTypes[j] {
				nx = x + dx[j]
				ny = y + dy[j]
			}
		}

		if nx < 1 || ny < 1 || nx > len(graph) || ny > len(graph) {
			continue
		}

		x = nx
		y = ny
	}

	writer.WriteString(strconv.Itoa(y))
	writer.WriteString(",")
	writer.WriteString(strconv.Itoa(x))
	writer.Flush()
}

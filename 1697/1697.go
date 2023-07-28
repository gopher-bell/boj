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
	line    = [100001]int{}
)

type q struct {
	idx  int
	dist int
}

func bfs(n1, n2 int) int {
	queue := []q{{n1, 1}}
	line[n1] = 1
	for len(queue) > 0 {
		dq := queue[0]
		queue = queue[1:]

		if dq.idx == n2 {
			return dq.dist
		}

		d1 := dq.idx - 1
		d2 := dq.idx + 1
		d3 := dq.idx * 2

		if d1 >= 0 && d1 < len(line) && line[d1] == 0 {
			line[d1] = dq.dist + 1
			queue = append(queue, q{d1, dq.dist + 1})
		}
		if d2 >= 0 && d2 < len(line) && line[d2] == 0 {
			line[d2] = dq.dist + 1
			queue = append(queue, q{d2, dq.dist + 1})
		}
		if d3 >= 0 && d3 < len(line) && line[d3] == 0 {
			line[d3] = dq.dist + 1
			queue = append(queue, q{d3, dq.dist + 1})
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

	res := bfs(n1, n2)

	writer.WriteString(strconv.Itoa(res - 1))
}

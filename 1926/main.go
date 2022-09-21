package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	graph := make([][]int, n)

	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, m)
		w := make([]int, m)
		for j := range w {
			fmt.Fscan(reader, &w[j])
		}
		graph[i] = w
	}

	fmt.Println(graph)
}

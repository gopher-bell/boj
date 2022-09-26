package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	colors = [2]string{"W", "B"}
	graph  []string
	y      int
	x      int
	minFix int
)

func paintGraph(y, x int) {
	defaultColor := y % 2
	var toChange int
	for j := y; j < y+8; j++ {
		for k := x; k < x+8; k++ {
			if (j+k)%2 != defaultColor {
				if string(graph[j][k]) == colors[defaultColor] {
					toChange++
				}
			} else {
				if string(graph[j][k]) != colors[defaultColor] {
					toChange++
				}
			}
		}
	}

	compare := 8*8 - toChange

	if toChange > compare {
		toChange = compare
	}

	if minFix > toChange {
		minFix = toChange
	}
}

func main() {
	defer writer.Flush()
	fmt.Fscanln(reader, &y, &x)
	graph = make([]string, y)
	for i := 0; i < y; i++ {
		fmt.Fscanf(reader, "%v\n", &graph[i])
	}

	minFix = y * x

	for i := 0; i < y-7; i++ {
		for j := 0; j < x-7; j++ {
			paintGraph(i, j)
		}
	}

	writer.WriteString(strconv.Itoa(minFix))
}

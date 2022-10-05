package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func readStdin() int {
	var s int
	if scanner.Scan() {
		s, _ = strconv.Atoi(scanner.Text())
	}

	return s
}

func getAboveAvg(keep []int, total int) float32 {
	avg := float32(total / len(keep))
	var n int
	for i := range keep {
		if float32(keep[i]) > avg {
			n++
		}
	}

	return float32(n) / float32(len(keep)) * 100
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := readStdin()

	for i := 0; i < n; i++ {
		num := readStdin()
		keep := make([]int, num)
		var total int
		for j := 0; j < num; j++ {
			s := readStdin()
			total += s
			keep[j] = s
		}
		res := getAboveAvg(keep, total)

		writer.WriteString(fmt.Sprintf("%.3f", res))
		writer.WriteString("%")
		writer.WriteString("\n")
	}
}

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

func getAboveAvg(keep []int, avg float32) int {
	var n int
	for i := range keep {
		if float32(keep[i]) > avg {
			n++
		}
	}
	return n
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
			keep[j] = s
			total += s
		}
		avg := float32(total / num)
		res := getAboveAvg(keep, avg)
		ans := float32(res % num)
		fmt.Println(res, num)
		writer.WriteString(fmt.Sprintf("%.3f", ans))
		writer.WriteString("%")
		writer.WriteString("\n")
	}
}

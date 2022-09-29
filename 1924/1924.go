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
	days    = []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}
	months  = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanLines)

	var n1, n2 int

	if scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		n1, _ = strconv.Atoi(s[0])
		n2, _ = strconv.Atoi(s[1])
	}

	var s int
	for i := 0; i < n1-1; i++ {
		s += months[i]
	}

	s = (s + n2) % 7

	writer.WriteString(days[s])

}

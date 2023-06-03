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
)

func main() {
	defer writer.Flush()
	scanner.Scan()
	s := strings.Fields(scanner.Text())
	s1, _ := strconv.Atoi(s[0])
	s2, _ := strconv.Atoi(s[1])

	ring := make([]int, s1)
	for i := 0; i < s1; i++ {
		ring[i] = i + 1
	}

	writer.WriteString("<")

	j := 0
	for i := 0; i < s1; i++ {
		if i > 0 {
			writer.WriteString(", ")
		}
		j = (j + s2 - 1) % len(ring)
		writer.WriteString(strconv.Itoa(ring[j]))
		ring = append(ring[:j], ring[j+1:]...)
	}

	writer.WriteString(">")
}

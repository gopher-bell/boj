package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func readStdin() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = readStdin()
	}

	sort.Slice(s, func(i, j int) bool {
		if len(s[i]) == len(s[j]) {
			return s[i] < s[j]
		} else {
			return len(s[i]) < len(s[j])
		}
	})

	for i := range s {
		if i > 0 && s[i-1] == s[i] {
			continue
		}
		writer.WriteString(s[i])
		writer.WriteString("\n")
	}
}

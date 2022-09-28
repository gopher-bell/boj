package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func readStdin() int {
	scanner.Scan()
	s, _ := strconv.Atoi(scanner.Text())
	return s
}

func max(s []int) int {
	var n int
	for i := range s {
		if s[i] > n {
			n = s[i]
		}
	}

	return n
}

func solution() int {
	nums := readStdin()
	nth := readStdin()
	s := make([]int, nums)

	for i := range s {
		s[i] = readStdin()
	}

	var cnt = 1

	for len(s) != 0 {
		maxNum := max(s)
		pop := s[0]
		s = s[1:]

		if pop == maxNum {
			if nth == 0 {
				break
			} else {
				nth--
				cnt++
			}
		} else {
			s = append(s, pop)
			nth--
			if nth == -1 {
				nth = len(s) - 1
			}
		}
	}

	return cnt
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	cases := readStdin()

	for i := 0; i < cases; i++ {
		s := solution()
		writer.WriteString(strconv.Itoa(s))
		writer.WriteString("\n")
	}

}

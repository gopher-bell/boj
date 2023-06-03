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

type stack struct {
	idx int
	val int
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	nums, _ := strconv.Atoi(scanner.Text())
	s := make([]stack, 0, nums)
	for i := 0; i < nums; i++ {
		scanner.Scan()
		tower, _ := strconv.Atoi(scanner.Text())
		flag := true
		for len(s) > 0 {
			if s[len(s)-1].val > tower {
				writer.WriteString(strconv.Itoa(s[len(s)-1].idx) + " ")
				s = append(s, stack{i + 1, tower})
				flag = false
				break
			} else {
				s = s[:len(s)-1]
			}
		}

		if flag {
			writer.WriteString("0 ")
			s = append(s, stack{i + 1, tower})
		}
	}
}

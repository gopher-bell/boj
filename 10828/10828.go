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
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	s := make([]int, 0)

	nums, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < nums; i++ {
		scanner.Scan()
		t := strings.Fields(scanner.Text())

		switch t[0] {
		case "size":
			writer.WriteString(strconv.Itoa(len(s)))
			writer.WriteString("\n")
		case "empty":
			if len(s) == 0 {
				writer.WriteString("1")
			} else {
				writer.WriteString("0")
			}
			writer.WriteString("\n")
		case "top":
			if len(s) == 0 {
				writer.WriteString("-1")
			} else {
				writer.WriteString(strconv.Itoa(s[len(s)-1]))
			}
			writer.WriteString("\n")
		case "push":
			n, _ := strconv.Atoi(t[1])
			s = append(s, n)
		default:
			if len(s) == 0 {
				writer.WriteString("-1")
			} else {
				x := s[len(s)-1]
				s = s[:len(s)-1]
				writer.WriteString(strconv.Itoa(x))
			}
			writer.WriteString("\n")
		}
	}
}

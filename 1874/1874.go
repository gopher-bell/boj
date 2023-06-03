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

func main() {
	defer writer.Flush()
	scanner.Scan()
	runs, _ := strconv.Atoi(scanner.Text())

	res := make([]string, 0, runs)
	stack := make([]int, 0, runs)
	push := 1
	for i := 0; i < runs; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		for push <= num {
			stack = append(stack, push)
			res = append(res, "+")
			push++
		}
		if stack[len(stack)-1] == num {
			stack = stack[:len(stack)-1]
			res = append(res, "-")
		} else {
			writer.WriteString("NO")
			return
		}
	}

	for _, v := range res {
		writer.WriteString(v)
		writer.WriteString("\n")
	}
}

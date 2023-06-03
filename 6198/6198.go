package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	writer  = bufio.NewWriter(os.Stdout)
	scanner = bufio.NewScanner(os.Stdin)
)

func getNum() int {
	scanner.Scan()
	b := scanner.Bytes()
	ret := 0
	for _, v := range b {
		ret *= 10
		ret += int(v - '0')
	}

	return ret
}

func main() {
	defer writer.Flush()
	nums := getNum()

	stack := make([]int, 0)
	buildings := make([]int, nums)
	for i := 0; i < nums; i++ {
		height := getNum()
		buildings[i] = height
	}

	cnt := 0
	for _, v := range buildings {
		for len(stack) > 0 && stack[len(stack)-1] <= v {
			stack = stack[:len(stack)-1]
		}
		cnt += len(stack)
		stack = append(stack, v)
	}

	writer.WriteString(strconv.Itoa(cnt))
}

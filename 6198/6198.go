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

func getNumber() int {
	ret := 0
	scanner.Scan()
	for _, v := range scanner.Bytes() {
		ret *= 10
		ret += int(v - '0')
	}

	return ret
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	runs := getNumber()

	stack := make([]int, 0, runs)
	cnt := 0
	for i := 0; i < runs; i++ {
		num := getNumber()
		for len(stack) > 0 && stack[len(stack)-1] <= num {
			stack = stack[:len(stack)-1]
		}

		cnt += len(stack)
		stack = append(stack, num)
	}

	writer.WriteString(strconv.Itoa(cnt))

}

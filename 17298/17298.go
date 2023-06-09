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

	numbers := make([]int, 0, runs)
	for i := 0; i < runs; i++ {
		num := getNumber()
		numbers = append(numbers, num)
	}

	stack := make([]int, 0, runs)
	for i, v := range numbers {
		for len(stack) > 0 && numbers[stack[len(stack)-1]] < v {
			numbers[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	for len(stack) > 0 {
		numbers[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}

	for _, v := range numbers {
		writer.WriteString(strconv.Itoa(v) + " ")
	}
}

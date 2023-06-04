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

func getNumbers() int {
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
	scanner.Split(bufio.ScanWords)
	run := getNumbers()
	nums := make([]int, 0, run)
	for i := 0; i < run; i++ {
		num := getNumbers()
		nums = append(nums, num)
	}

	stack := make([]int, 0, run)
	for i := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			nums[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	for len(stack) > 0 {
		nums[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}

	for i := range nums {
		writer.WriteString(strconv.Itoa(nums[i]))
		if i != len(nums)-1 {
			writer.WriteString(" ")
		}
	}
}

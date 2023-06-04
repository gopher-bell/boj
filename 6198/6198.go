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

type tower struct {
	idx    int
	height int
}

func getNumber() int {
	scanner.Scan()
	b := scanner.Bytes()
	ans := 0

	for _, v := range b {
		ans *= 10
		ans += int(v - '0')
	}

	return ans
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	nums := getNumber()
	cnt := 0
	stack := make([]int, 0, nums)
	for i := 0; i < nums; i++ {
		height := getNumber()
		flag := true
		for len(stack) > 0 {
			last := len(stack) - 1
			if stack[last] > height {
				cnt += len(stack)
				stack = append(stack, height)
				flag = false
				break
			} else {
				stack = stack[:last]
			}
		}

		if flag {
			stack = append(stack, height)
		}
	}

	writer.WriteString(strconv.Itoa(cnt))
}

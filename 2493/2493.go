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
	tmp := 0

	for _, v := range b {
		ans *= 10
		tmp = int(v - '0')
		ans += tmp
		tmp = 0
	}

	return ans
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	nums := getNumber()
	stack := make([]tower, 0, nums)
	for i := 0; i < nums; i++ {
		height := getNumber()
		flag := true
		for len(stack) > 0 {
			last := len(stack) - 1
			if stack[last].height >= height {
				writer.WriteString(strconv.Itoa(stack[last].idx) + " ")
				stack = append(stack, tower{idx: i + 1, height: height})
				flag = false
				break
			} else {
				stack = stack[:last]
			}
		}

		if flag {
			stack = append(stack, tower{idx: i + 1, height: height})
			writer.WriteString("0 ")
		}
	}
}

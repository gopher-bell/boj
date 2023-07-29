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
	flag := false
	scanner.Scan()
	for _, v := range scanner.Bytes() {
		if v == '-' {
			flag = true
			continue
		}

		ret *= 10
		ret += int(v - '0')
	}

	if flag {
		return -ret
	}

	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n1, n2 := getNumber(), getNumber()
	nums := make([]int, n1+1)
	sums := make([]int, n1+1)
	for i := 1; i < len(nums); i++ {
		nums[i] = getNumber()
		if i == 1 {
			sums[i] = nums[i]
		} else {
			sums[i] = sums[i-1] + nums[i]
		}
	}

	m := -1000000000
	for i := n2; i < len(nums); i++ {
		m = max(m, sums[i]-sums[i-n2])
	}

	writer.WriteString(strconv.Itoa(m))
}

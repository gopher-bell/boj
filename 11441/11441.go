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

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n1 := getNumber()
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

	n2 := getNumber()

	for i := 0; i < n2; i++ {
		t1, t2 := getNumber(), getNumber()
		res := sums[t2] - sums[t1-1]
		writer.WriteString(strconv.Itoa(res))
		if i != n2-1 {
			writer.WriteString("\n")
		}
	}
}

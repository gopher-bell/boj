package main

import (
	"bufio"
	"os"
	"sort"
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
	tests := getNumber()
	nums := make([]int, tests)
	for i := 0; i < tests; i++ {
		nums[i] = getNumber()
	}

	sort.Ints(nums)

	for i, v := range nums {
		writer.WriteString(strconv.Itoa(v))
		if i != len(nums)-1 {
			writer.WriteString("\n")
		}
	}
}

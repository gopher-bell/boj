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
	nums    []int
	arr     []int
	visited []bool
	n, m    int
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

func dfs(k int) {
	if k == m {
		for i := 0; i < m; i++ {
			writer.WriteString(strconv.Itoa(arr[i]))
			writer.WriteString(" ")
		}
		writer.WriteString("\n")

		return
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			arr[k] = nums[i]
			visited[i] = true
			dfs(k + 1)
			visited[i] = false
		}
	}
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = getNumber(), getNumber()
	nums = make([]int, n)
	arr = make([]int, n)
	visited = make([]bool, n)
	for i := 0; i < n; i++ {
		nums[i] = getNumber()
	}
	sort.Ints(nums)
	dfs(0)
}

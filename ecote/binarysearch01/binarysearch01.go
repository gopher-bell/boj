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

func getWord() int {
	scanner.Scan()
	s, _ := strconv.Atoi(scanner.Text())
	return s
}

func binarySearch(nums []int, start, end, weight int) int {
	var result int
	for start <= end {
		mid := (start + end) / 2
		var total int
		for i := range nums {
			if mid < nums[i] {
				total += nums[i] - mid
			}
		}
		if total < weight {
			end = mid - 1
		} else {
			result = mid
			start = mid + 1
		}
	}

	return result
}

func main() {
	scanner.Split(bufio.ScanWords)
	n1, n2 := getWord(), getWord()
	inputs := make([]int, n1)

	for i := range inputs {
		inputs[i] = getWord()
	}

	sort.Ints(inputs)

	s := binarySearch(inputs, 0, inputs[len(inputs)-1], n2)

	writer.WriteString(strconv.Itoa(s))
	writer.Flush()
}

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

func readWord() int {
	if scanner.Scan() {
		s, _ := strconv.Atoi(scanner.Text())
		return s
	}
	return -1
}

func binarySearch(nums []int, find int) bool {
	var start int
	var end int

	start = 0
	end = len(nums) - 1

	for start <= end {
		mid := (start + end) / 2
		if nums[mid] > find {
			end = mid - 1
		} else if nums[mid] < find {
			start = mid + 1
		} else {
			return true
		}
	}
	return false
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n1 := readWord()
	nums1 := make([]int, n1)
	for i := 0; i < n1; i++ {
		nums1[i] = readWord()
	}
	n2 := readWord()
	nums2 := make([]int, n2)
	for i := 0; i < n2; i++ {
		nums2[i] = readWord()
	}

	sort.Ints(nums1)

	for i := range nums2 {
		if binarySearch(nums1, nums2[i]) {
			writer.WriteString("1")
			writer.WriteString("\n")
		} else {
			writer.WriteString("0")
			writer.WriteString("\n")
		}
	}
}

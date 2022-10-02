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

	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})

}

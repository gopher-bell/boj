package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	result  = 0
)

func getWords() []string {
	scanner.Scan()
	return strings.Fields(scanner.Text())
}

func binarySearch(nums []int, start, end, weight int) {
	var total int
	if start <= end {
		mid := (start + end) / 2
		fmt.Println(start, end, mid)

		for i := range nums {
			if nums[i] > mid {
				total += nums[i] - mid
			}
		}
		if total < weight {
			binarySearch(nums, start, mid-1, weight)
		} else {
			result = mid
			binarySearch(nums, mid+1, end, weight)
		}
	}
}

func main() {
	s := getWords()

	n1, _ := strconv.Atoi(s[0])
	n2, _ := strconv.Atoi(s[1])

	slice := make([]int, n1)

	t := getWords()

	for i := range slice {
		slice[i], _ = strconv.Atoi(t[i])
	}

	sort.Ints(slice)

	binarySearch(slice, 0, slice[len(slice)-1], n2)

	fmt.Println(result)
}

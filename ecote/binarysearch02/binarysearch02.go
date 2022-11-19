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

func getWord() int {
	scanner.Scan()
	s, _ := strconv.Atoi(scanner.Text())
	return s
}

func binarySearchFirst(nums []int, start, end, k int) int {
	if start <= end {
		mid := (start + end) / 2
		if nums[mid] == k {
			if mid-1 >= 0 && nums[mid-1] == k {
				return binarySearchFirst(nums, start, mid-1, k)
			} else {
				return mid
			}
		} else if nums[mid] > k {
			return binarySearchFirst(nums, start, mid-1, k)
		} else {
			return binarySearchFirst(nums, mid+1, end, k)
		}
	}

	return -1
}

func binarySearchLast(nums []int, start, end, k int) int {
	if start <= end {
		mid := (start + end) / 2
		if nums[mid] == k {
			if mid+1 < len(nums) && nums[mid+1] == k {
				return binarySearchLast(nums, mid+1, end, k)
			} else {
				return mid
			}
		} else if nums[mid] > k {
			return binarySearchLast(nums, start, mid-1, k)
		} else {
			return binarySearchLast(nums, mid+1, end, k)
		}
	}

	return -1
}

func main() {
	scanner.Split(bufio.ScanWords)
	n1, n2 := getWord(), getWord()

	input := make([]int, n1)
	for i := range input {
		input[i] = getWord()
	}

	res := binarySearchFirst(input, 0, len(input)-1, n2)
	res2 := binarySearchLast(input, 0, len(input)-1, n2)

	writer.WriteString(strconv.Itoa(res) + "\n")
	writer.WriteString(strconv.Itoa(res2))
	writer.Flush()
}

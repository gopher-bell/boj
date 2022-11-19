package main

import "fmt"

func quickSort(nums []int, start, end int) {
	if start < end {
		p := partition(nums, start, end)
		quickSort(nums, start, p-1)
		quickSort(nums, p+1, end)
	}
}

func partition(nums []int, start, end int) int {
	pivot := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < pivot {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]

	return i
}

func main() {
	s := []int{1, 6, 10, 9, 2, 7, 8, 3, 5, 4}
	quickSort(s, 0, len(s)-1)
	fmt.Println(s)
}

package main

import "fmt"

func selectionSort(nums []int) {
	for i := range nums {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		nums[min], nums[i] = nums[i], nums[min]
	}
}

func main() {
	s := []int{1, 6, 10, 9, 2, 7, 8, 3, 5, 4}
	selectionSort(s)
	fmt.Println(s)
}

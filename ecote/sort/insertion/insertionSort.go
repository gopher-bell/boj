package main

import "fmt"

func insertionSort(nums []int) {
	for i := range nums {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				fmt.Println(nums)
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				break
			}
		}
	}
}

func main() {
	s := []int{1, 6, 10, 9, 2, 7, 8, 3, 5, 4}
	insertionSort(s)
	fmt.Println(s)
}

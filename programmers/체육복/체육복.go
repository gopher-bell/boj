package main

import "fmt"

func solution(n int, lost []int, reserve []int) int {
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = 1
	}
	for _, val := range lost {
		a[val]--
	}
	for _, val := range reserve {
		a[val]++
	}
	fmt.Println(a)
	ans := 0
	for i := 1; i <= n; i++ {
		if a[i] == 0 {
			if a[i-1] > 1 {
				a[i-1]--
				a[i]++
			} else if i < n && a[i+1] > 1 {
				a[i+1]--
				a[i]++
			}
		}
		if a[i] > 0 {
			ans++
		}
	}
	return ans
}

func main() {
	res := solution(13, []int{1, 2, 5, 6, 10, 12, 13}, []int{2, 3, 4, 5, 7, 8, 9, 10, 11, 12})
	fmt.Println(res)
}

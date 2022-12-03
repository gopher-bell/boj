package main

import (
	"fmt"
	"math"
)

func solution(number int, limit int, power int) int {
	div := make([]int, number)
	for i := 1; i <= number; i++ {
		cnt := 0
		for j := 1; j <= int(math.Sqrt(float64(i))); j++ {
			fmt.Println(i, j)
			if i%j == 0 {
				if i/j == j {
					cnt++
				} else {
					cnt += 2
				}
			}
		}
		div[i-1] = cnt
	}

	fmt.Println(div)

	weight := 0
	for i := range div {
		if div[i] <= limit {
			weight += div[i]
		} else {
			weight += power
		}
	}

	return weight
}

func main() {
	res := solution(100, 3, 2)
	fmt.Println(res)
}

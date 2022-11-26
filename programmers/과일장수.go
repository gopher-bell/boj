package main

import "sort"

func solution(k int, m int, score []int) int {
	var tmp []int
	total := 0
	sort.Ints(score)
	for i := len(score) - 1; i >= 0; i-- {
		if score[i] > k {
			continue
		}
		tmp = append(tmp, score[i])
		if len(tmp) == m {
			min := tmp[0]
			for j := 1; j < len(tmp); j++ {
				if tmp[j] < min {
					min = tmp[j]
				}
			}
			tmp = nil
			total += min * m
		}
	}

	return total
}

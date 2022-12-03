package 명예의전당

import "sort"

func solution(k int, score []int) []int {
	var res []int
	var ans []int
	for i := range score {
		res = append(res, score[i])
		sort.Slice(res, func(i, j int) bool {
			return res[i] < res[j]
		})
		if len(res) > k {
			res = res[1:]
		}
		ans = append(ans, res[0])
	}

	return ans
}

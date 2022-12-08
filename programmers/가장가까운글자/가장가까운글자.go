package 가장가까운글자

func solution(s string) []int {
	m := make(map[string]int)
	var res []int

	for i := range s {
		val, ok := m[string(s[i])]
		if !ok {
			m[string(s[i])] = i
			res = append(res, -1)
		} else {
			m[string(s[i])] = i
			res = append(res, i-val)
		}
	}
	return res
}

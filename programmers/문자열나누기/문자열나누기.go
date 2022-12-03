package 문자열나누기

func solution(s string) int {
	comp := s[0]
	cnt1 := 0
	cnt2 := 0
	res := 0
	for i := 0; i < len(s); i++ {
		if comp == s[i] {
			cnt1++
		} else {
			cnt2++
		}
		if cnt1 == cnt2 {
			res++
			if i+1 >= len(s) {
				return res
			} else {
				comp = s[i+1]
			}
		}
	}

	return res + 1
}

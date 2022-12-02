package 문자열나누기

func solution(s string) int {
	x := s[0]
	cnt := 0
	cnt1 := 0
	cnt2 := 0
	for idx, item := range s {
		if item == rune(x) {
			cnt1 += 1
		} else {
			cnt2 += 1
		}
		if cnt1 == cnt2 {
			cnt += 1
			if idx+1 != len(s) {
				x = s[idx+1]
			} else {
				return cnt
			}
		}
	}
	return cnt + 1
}

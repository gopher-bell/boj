package 약수의합

func solution(n int) int {
	answer := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			answer += i
		}
	}

	return answer
}

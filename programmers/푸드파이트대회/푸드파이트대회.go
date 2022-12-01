package 푸드파이트대회

import (
	"bytes"
	"strconv"
)

func solution(food []int) string {
	b := bytes.Buffer{}
	for i := range food {
		if food[i] < 2 {
			continue
		}
		for j := 0; j < food[i]/2; j++ {
			b.WriteString(strconv.Itoa(i))
		}
	}

	b.WriteString("0")

	for i := len(food) - 1; i >= 0; i-- {
		if food[i] < 2 {
			continue
		}
		for j := 0; j < food[i]/2; j++ {
			b.WriteString(strconv.Itoa(i))
		}
	}

	return b.String()
}

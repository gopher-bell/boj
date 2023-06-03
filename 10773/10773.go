package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func getNum(b []byte) int {
	ans := 0
	tmp := 0
	for i := range b {
		if b[0] == '0' {
			break
		}
		if b[i] >= '0' && '9' >= b[i] {
			tmp *= 10
			tmp += int(b[i] - '0')
			continue
		}

		ans += tmp
		tmp = 0
	}

	return ans + tmp
}

func main() {
	defer writer.Flush()

	b1, _ := reader.ReadBytes('\n')
	totalRuns := getNum(b1)

	arr := make([]int, 0, totalRuns)
	for i := 0; i < totalRuns; i++ {
		b2, _ := reader.ReadBytes('\n')
		num := getNum(b2)
		if num != 0 {
			arr = append(arr, num)
		} else {
			if len(arr) > 0 {
				arr = arr[:len(arr)-1]
			}
		}
	}

	max := 0
	for _, v := range arr {
		max += v
	}

	writer.WriteString(strconv.Itoa(max))
}

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	n       int
	total   int
	closest int
)

func readInput() []int {
	scanner.Scan()
	s := strings.Fields(scanner.Text())
	res := make([]int, n)
	for i := range res {
		res[i], _ = strconv.Atoi(s[i])
	}

	return res
}

func main() {
	defer writer.Flush()
	scanner.Scan()
	s := strings.Fields(scanner.Text())
	n, _ = strconv.Atoi(s[0])
	total, _ = strconv.Atoi(s[1])
	nums := readInput()

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				current := nums[i] + nums[j] + nums[k]
				if total >= current {
					if current > closest {
						closest = current
					}
				}
			}
		}
	}

	writer.WriteString(strconv.Itoa(closest))
}

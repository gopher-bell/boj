package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	writer  = bufio.NewWriter(os.Stdout)
	scanner = bufio.NewScanner(os.Stdin)
)

func main() {
	defer writer.Flush()
	scanner.Scan()
	s1 := strings.Split(scanner.Text(), " ")
	n1, _ := strconv.Atoi(s1[0])
	n2, _ := strconv.Atoi(s1[1])

	dp := make([][]int, n1+1)
	graph := make([][]int, n1+1)
	for i := range graph {
		dp[i] = make([]int, n1+1)
		graph[i] = make([]int, n1+1)
		if i > 0 {
			scanner.Scan()
			s2 := strings.Split(scanner.Text(), " ")
			for j := 1; j < len(graph[i]); j++ {
				num, _ := strconv.Atoi(s2[j-1])
				graph[i][j] = num
				dp[i][j] = num
			}
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + graph[i][j]
		}
	}

	for i := 0; i < n2; i++ {
		scanner.Scan()
		s3 := strings.Split(scanner.Text(), " ")
		n1, _ := strconv.Atoi(s3[0])
		n2, _ := strconv.Atoi(s3[1])
		n3, _ := strconv.Atoi(s3[2])
		n4, _ := strconv.Atoi(s3[3])

		res := dp[n3][n4] - dp[n3][n2-1] - dp[n1-1][n4] + dp[n1-1][n2-1]
		writer.WriteString(strconv.Itoa(res) + "\n")
	}
}

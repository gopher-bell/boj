package main

import "fmt"

type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	s := make([][]int, len(matrix)+1)
	for i := range s {
		s[i] = make([]int, len(matrix[0])+1)
		if i > 0 {
			for j := 1; j < len(s[i]); j++ {
				s[i][j] = s[i][j-1] + s[i-1][j] - s[i-1][j-1] + matrix[i-1][j-1]
			}
		}
	}

	n := NumMatrix{
		sum: s,
	}

	return n
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sum[row2+1][col2+1] - this.sum[row1][col2+1] - this.sum[row2+1][col1] + this.sum[row1][col1]
}

func main() {
	ss := [][]int{
		[]int{3, 0, 1, 4, 2},
		[]int{5, 6, 3, 2, 1},
		[]int{1, 2, 0, 1, 5},
		[]int{4, 1, 0, 1, 7},
		[]int{1, 0, 3, 0, 5},
	}

	res := Constructor(ss)

	for _, v := range res.sum {
		fmt.Println(v)
	}
}

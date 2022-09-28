package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	res     int
	need    int
)

func readStdin() int {
	scanner.Scan()
	s, _ := strconv.Atoi(scanner.Text())
	return s
}

func solution(trees []int) {
	min := 0
	max := trees[len(trees)-1]

	for min <= max {
		var total int
		mid := (min + max) / 2

		for i := range trees {
			if trees[i] > mid {
				total += trees[i] - mid
			}
		}

		if total >= need {
			res = mid
			min = mid + 1
		}
		if need > total {
			max = mid - 1
		}
	}
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	numOfTrees := readStdin()
	need = readStdin()
	trees := make([]int, numOfTrees)
	for i := 0; i < numOfTrees; i++ {
		trees[i] = readStdin()
	}
	sort.Slice(trees, func(i, j int) bool {
		return trees[i] < trees[j]
	})

	solution(trees)

	writer.WriteString(strconv.Itoa(res))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var text1, text2 string
	fmt.Fscan(reader, &text1, &text2)

	arr1 := make([]int, 26)
	arr2 := make([]int, 26)

	for i := range text1 {
		arr1[text1[i]-97]++
	}
	for i := range text2 {
		arr2[text2[i]-97]++
	}

	cnt := 0
	for i := 0; i < 26; i++ {
		if arr1[i] != arr2[i] {
			cnt += abs(arr1[i], arr2[i])
		}
	}

	writer.WriteString(strconv.Itoa(cnt))
}

func abs(x, y int) int {
	if x < y {
		return y - x
	}

	return x - y
}

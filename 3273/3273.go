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

	nums := 0
	fmt.Fscan(reader, &nums)

	arr := make([]int, 0, nums)
	for i := 0; i < nums; i++ {
		x := 0
		fmt.Fscan(reader, &x)

		arr = append(arr, x)
	}

	tot := 0
	fmt.Fscan(reader, &tot)

	res := make([]int, tot)
	cnt := 0
	for _, v := range arr {
		if v >= tot {
			continue
		}

		if res[tot-v] == 1 {
			cnt++
		}

		res[v]++
	}

	writer.WriteString(strconv.Itoa(cnt))
}

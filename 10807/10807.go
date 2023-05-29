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
	num := 0
	fmt.Fscan(reader, &num)
	arr := make([]int, 0, num)
	for i := 0; i < num; i++ {
		x := 0
		fmt.Fscan(reader, &x)
		arr = append(arr, x)
	}

	find := 0
	fmt.Fscan(reader, &find)

	cnt := 0
	for _, v := range arr {
		if v == find {
			cnt++
		}
	}

	writer.WriteString(strconv.Itoa(cnt))
}

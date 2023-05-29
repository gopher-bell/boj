package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	text := scanner.Text()
	arr := make([]int, 11)

	for i := range text {
		if text[i]-48 == 6 || text[i]-48 == 9 {
			if arr[6] > arr[9] {
				arr[9]++
			} else {
				arr[6]++
			}

			continue
		}

		arr[text[i]-48]++
	}

	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	writer.WriteString(strconv.Itoa(max))
}

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
	n, _ := strconv.Atoi(scanner.Text())
	var cnt int
	for {
		if n%5 != 0 {
			if n < 3 {
				cnt = -1
				break
			}
			n = n - 3
			cnt++
		} else {
			break
		}
	}
	if cnt != -1 {
		cnt += n / 5
	}

	writer.WriteString(strconv.Itoa(cnt))
}

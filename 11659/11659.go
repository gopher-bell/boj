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

func getNumber() int {
	ret := 0
	scanner.Scan()
	for _, v := range scanner.Bytes() {
		ret *= 10
		ret += int(v - '0')
	}

	return ret
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n1, n2 := getNumber(), getNumber()
	s := make([]int, n1+1)
	r := make([]int, n1+1)
	for i := 1; i < len(s); i++ {
		s[i] = getNumber()
		if i == 1 {
			r[i] = s[i]
		} else {
			r[i] = r[i-1] + s[i]
		}
	}

	for i := 0; i < n2; i++ {
		t1, t2 := getNumber(), getNumber()
		res := r[t2] - r[t1-1]
		writer.WriteString(strconv.Itoa(res))
		if i != n2-1 {
			writer.WriteString("\n")
		}
	}
}

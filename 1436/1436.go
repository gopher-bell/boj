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
	num := getNumber()
	cnt := 1
	s := 666
	for num != cnt {
		s++
		if strings.Contains(strconv.Itoa(s), "666") {
			cnt++
		}
	}

	writer.WriteString(strconv.Itoa(s))
}

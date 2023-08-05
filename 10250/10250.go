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
	t := getNumber()
	for i := 0; i < t; i++ {
		h, _, n := getNumber(), getNumber(), getNumber()

		if n%h == 0 {
			writer.WriteString(strconv.Itoa((h*100)+(n/h)) + "\n")
		} else {
			writer.WriteString(strconv.Itoa(((n%h)*100)+((n/h)+1)) + "\n")
		}
	}
}

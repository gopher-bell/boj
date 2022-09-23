package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	s, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	s = strings.TrimSpace(s)

	var cnt int

	t := strings.Split(s, " ")
	for i := range t {
		if t[i] != "" {
			cnt++
		}
	}

	writer.WriteString(strconv.Itoa(cnt))
}

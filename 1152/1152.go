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
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	res := strings.Fields(s)
	writer.WriteString(strconv.Itoa(len(res)))
}

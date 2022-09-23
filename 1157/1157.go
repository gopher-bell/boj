package main

import (
	"bufio"
	"os"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	s, _ := reader.ReadString('\n')
	s = strings.ToUpper(strings.TrimSpace(s))
	var box [26]int
	for i := range s {
		if s[i] >= 'A' && 'Z' >= s[i] {
			box[s[i]-'A']++
		}
	}
	var maxInt = -1
	var maxKey string
	for i := range box {
		if box[i] > maxInt {
			if !(box[i] == 0) {
				maxInt = box[i]
				maxKey = string(i + 'A')
			}
		} else if box[i] == maxInt {
			maxKey = "?"
		}
	}

	writer.WriteString(maxKey)
}

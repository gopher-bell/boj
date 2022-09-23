package main

import (
	"bufio"
	"os"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	s := strings.ToLower(scanner.Text())
	var box [26]int
	for i := range s {
		if s[i] >= 'a' && 'z' >= s[i] {
			box[s[i]-'a']++
		}
	}

	var maxInt = -1
	var maxKey string
	for i := range box {
		if box[i] > maxInt {
			maxInt = box[i]
			maxKey = strings.ToUpper(string(rune(i) + 'a'))
		} else if box[i] == maxInt {
			maxKey = "?"
		}
	}

	writer.WriteString(maxKey)
}

package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func getWords() ([]string, int) {
	scanner.Scan()
	s := strings.Split(scanner.Text(), "")
	var words []string
	var num int
	for _, v := range s {
		input, err := strconv.Atoi(v)
		if err != nil {
			words = append(words, v)
			continue
		}
		num += input
	}

	return words, num
}

func main() {
	words, nums := getWords()
	sort.Strings(words)

	for _, v := range words {
		writer.WriteString(v)
	}
	writer.WriteString(strconv.Itoa(nums))
	writer.Flush()
}

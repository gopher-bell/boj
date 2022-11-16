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

func getWords() []int {
	scanner.Scan()
	s := strings.Split(scanner.Text(), " ")
	res := make([]int, len(s))
	for i, v := range s {
		res[i], _ = strconv.Atoi(v)
	}
	return res
}

func main() {
	fear := getWords()
	sort.Ints(fear)
	var grp, cnt int

	for _, v := range fear {
		cnt++
		if v <= cnt {
			grp++
			cnt = 0
		}
	}

	writer.WriteString(strconv.Itoa(grp))
	writer.Flush()
}

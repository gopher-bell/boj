package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()

	var text string
	var n int
	fmt.Fscanln(reader, &text)
	fmt.Fscanln(reader, &n)

	l := list.New()
	r := list.New()

	for _, v := range text {
		l.PushBack(string(v))
	}

	for i := 0; i < n; i++ {
		text, _ = reader.ReadString('\n')
		strings.TrimSuffix(text, "\n")
		switch string(text[0]) {
		case "L":
			ele := l.Back()
			if ele != nil {
				l.Remove(ele)
				r.PushFront(ele.Value)
			}
		case "D":
			ele := r.Front()
			if ele != nil {
				r.Remove(ele)
				l.PushBack(ele.Value)
			}
		case "B":
			ele := l.Back()
			if ele != nil {
				l.Remove(ele)
			}
		case "P":
			l.PushBack(string(text[2]))
		}
	}

	for e := l.Front(); e != nil; e = e.Next() {
		writer.WriteString(fmt.Sprint(e.Value))
	}

	for e := r.Front(); e != nil; e = e.Next() {
		writer.WriteString(fmt.Sprint(e.Value))
	}
}

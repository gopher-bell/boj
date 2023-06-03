package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var num int
	fmt.Fscanln(reader, &num)

	for i := 0; i < num; i++ {
		left := make([]byte, 0)
		right := make([]byte, 0)
		b, _ := reader.ReadBytes('\n')
		for _, v := range b {
			switch v {
			case '<':
				if len(left) > 0 {
					right = append(right, left[len(left)-1])
					left = left[:len(left)-1]
				}
			case '>':
				if len(right) > 0 {
					left = append(left, right[len(right)-1])
					right = right[:len(right)-1]
				}
			case '-':
				if len(left) > 0 {
					left = left[:len(left)-1]
				}
			default:
				if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9') {
					left = append(left, v)
				}
			}
		}

		for len(right) > 0 {
			left = append(left, right[len(right)-1])
			right = right[:len(right)-1]
		}

		for _, v := range left {
			writer.WriteByte(v)
		}

		writer.WriteByte('\n')
	}
}

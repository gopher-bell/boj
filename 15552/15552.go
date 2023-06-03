package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	n, _, _ := reader.ReadLine()
	nums := 0
	for _, v := range n {
		nums *= 10
		nums += int(v - '0')
	}

	for i := 0; i < nums; i++ {
		b, _, _ := reader.ReadLine()
		ans := 0
		tmp := 0
		for _, v := range b {
			if v == ' ' {
				ans = tmp
				tmp = 0
			} else {
				tmp *= 10
				tmp += int(v - '0')
			}
		}

		writer.WriteString(strconv.Itoa(ans + tmp))
		writer.WriteString("\n")
	}
}

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
	num := 0
	fmt.Fscan(reader, &num)

	for i := 0; i < num; i++ {
		var str1, str2 string
		fmt.Fscan(reader, &str1, &str2)

		if len(str1) != len(str2) {
			writer.WriteString("Impossible")
			writer.WriteString("\n")
			continue
		}

		arr1 := make([]int, 26)
		for j := range str1 {
			arr1[str1[j]-'a']++
			arr1[str2[j]-'a']--
		}

		flag := true
		for l := 0; l < 26; l++ {
			if arr1[l] != 0 {
				flag = false
				break
			}
		}

		if flag {
			writer.WriteString("Possible")
			writer.WriteString("\n")
		} else {
			writer.WriteString("Impossible")
			writer.WriteString("\n")
		}
	}
}

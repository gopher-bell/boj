package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func getWord() int {
	scanner.Scan()
	r, _ := strconv.Atoi(scanner.Text())
	return r
}

func main() {
	n1, n2 := getWord(), getWord()
	var count int

	for {
		//if n1 == 1 {
		//	break
		//} else if n1%n2 == 0 {
		//	n1 /= n2
		//	count++
		//} else {
		//	n1--
		//	count++
		//}
		fmt.Println("n1 1", n1)
		target := (n1 / n2) * n2
		fmt.Println("target", target)
		count += n1 - target
		fmt.Println("count", count)
		n1 = target
		fmt.Println("n1 2", n1)
		fmt.Println("-----")
		if n1 < n2 {
			break
		}
		count += 1
		n1 /= n2
	}

	count += n1 - 1

	writer.WriteString(strconv.Itoa(count))
	writer.Flush()
}

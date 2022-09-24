package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	total   float32
	max     float32
	avg     float32
)

type maxHeap []float32

func (n *maxHeap) Len() int {
	return len(*n)
}

func (n *maxHeap) Less(i, j int) bool {
	return (*n)[j] < (*n)[i]
}

func (n *maxHeap) Swap(i, j int) {
	(*n)[i], (*n)[j] = (*n)[j], (*n)[i]
}

func (n *maxHeap) Push(x interface{}) {
	*n = append(*n, x.(float32))
}

func (n *maxHeap) Pop() interface{} {
	old := *n
	l := len(old)
	x := old[l-1]
	*n = old[0 : l-1]
	return x
}

func scanStdin(n int) maxHeap {
	res := make(maxHeap, 0)
	heap.Init(&res)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		f, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return nil
		}
		heap.Push(&res, float32(f))
	}

	return res
}

func addToTotal(n float32) {
	n = n / max * 100
	total += n

}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	s, _ := strconv.Atoi(scanner.Text())
	scores := scanStdin(s)
	currLen := scores.Len()

	max = heap.Pop(&scores).(float32)
	addToTotal(max)

	for scores.Len() != 0 {
		m := heap.Pop(&scores).(float32)
		addToTotal(m)
	}

	avg = total / float32(currLen)

	writer.WriteString(fmt.Sprintf("%f", avg))
}

package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

type node struct {
	weight int
	index  int
}

type nodeHeap []*node

func (h nodeHeap) Len() int           { return len(h) }
func (h nodeHeap) Less(i, j int) bool { return h[i].weight < h[j].weight }
func (h nodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *nodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*node))
}

func (h *nodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getWords() []int {
	scanner.Scan()
	s := strings.Fields(scanner.Text())
	res := make([]int, len(s))
	for i := range s {
		res[i], _ = strconv.Atoi(s[i])
	}

	return res
}

func dijkstra(nums [][]int, start int) []int {
	visited := make([]bool, len(nums))
	distance := make([]int, len(nums))

	for i := range distance {
		distance[i] = 1000000000
	}
	distance[start] = 0

	h := &nodeHeap{&node{index: start, weight: 0}}
	heap.Init(h)

	for h.Len() > 0 {
		n := heap.Pop(h).(*node)
		visited[n.index] = true

		for i := range nums[n.index] {
			if nums[n.index][i] != 0 && !visited[i] {
				heap.Push(h, &node{index: i, weight: nums[n.index][i]})
				if distance[n.index]+nums[n.index][i] < distance[i] {
					distance[i] = distance[n.index] + nums[n.index][i]
				}
			}
		}
	}

	return distance
}

func main() {
	n := getWords()
	start := getWords()

	nums := make([][]int, n[0]+1)
	for i := range nums {
		nums[i] = make([]int, n[0]+1)
	}

	for i := 0; i < n[1]; i++ {
		s := getWords()
		nums[s[0]][s[1]] = s[2]
	}

	res := dijkstra(nums, start[0])

	for i := 1; i < len(res); i++ {
		if res[i] == 1000000000 {
			writer.WriteString("INF" + "" + "\n")
		} else {
			writer.WriteString(strconv.Itoa(res[i]) + "\n")
		}
	}
	writer.Flush()
}

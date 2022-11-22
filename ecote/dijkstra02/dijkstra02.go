package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

type node struct {
	index  int
	weight int
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

func getWord() int {
	scanner.Scan()
	s, _ := strconv.Atoi(scanner.Text())
	return s
}

func dijkstra(nums map[int][]node, visited []bool, distance []int, start int) []int {
	for i := range distance {
		distance[i] = 1000000000
	}
	distance[start] = 0

	h := &nodeHeap{&node{index: start, weight: 0}}
	heap.Init(h)

	for h.Len() > 0 {
		n := heap.Pop(h).(*node)
		if visited[n.index] {
			continue
		}
		visited[n.index] = true
		for _, v := range nums[n.index] {
			if !visited[v.index] {
				if distance[n.index]+v.weight < distance[v.index] {
					distance[v.index] = distance[n.index] + v.weight
					heap.Push(h, &node{index: v.index, weight: distance[v.index]})
				}
			}
		}
	}

	return distance
}

func main() {
	scanner.Split(bufio.ScanWords)

	nodes, edges, start := getWord(), getWord(), getWord()

	nums := make(map[int][]node)

	for i := 0; i < edges; i++ {
		from, to, weight := getWord(), getWord(), getWord()
		nums[from] = append(nums[from], node{to, weight})
	}

	visited := make([]bool, nodes+1)
	distance := make([]int, nodes+1)

	res := dijkstra(nums, visited, distance, start)

	for i := 1; i < len(res); i++ {
		if res[i] == 1000000000 {
			writer.WriteString("INF" + "" + "\n")
		} else {
			writer.WriteString(strconv.Itoa(res[i]) + "\n")
		}
	}
	writer.Flush()
}

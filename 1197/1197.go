package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
)

var (
	scanner     = bufio.NewScanner(os.Stdin)
	writer      = bufio.NewWriter(os.Stdout)
	totalWeight int
)

type Graph struct {
	Nodes []int
	Edges map[int][]Node
}

type Node struct {
	Index int
	Value int
}

type NodeHeap []*Node

func (n NodeHeap) Len() int {
	return len(n)
}

func (n NodeHeap) Less(i, j int) bool {
	return n[i].Value < n[j].Value
}

func (n NodeHeap) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n *NodeHeap) Push(x interface{}) {
	*n = append(*n, x.(*Node))
}

func (n *NodeHeap) Pop() interface{} {
	old := *n
	i := len(old)
	x := old[i-1]
	*n = old[0 : i-1]
	return x
}

func readStdin() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func (g *Graph) prim() {
	visited := make([]bool, len(g.Nodes))
	nodeHeap := NodeHeap{&Node{Index: 1, Value: 0}}
	heap.Init(&nodeHeap)
	for nodeHeap.Len() != 0 {
		n := heap.Pop(&nodeHeap).(*Node)
		if !visited[n.Index] {
			visited[n.Index] = true
			totalWeight += n.Value
			for i := range g.Edges[n.Index] {
				if !visited[g.Edges[n.Index][i].Index] {
					heap.Push(&nodeHeap, &(g.Edges[n.Index][i]))
				}
			}
		}
	}
}

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	v := readStdin()
	e := readStdin()

	graph := Graph{
		Nodes: make([]int, v+1),
		Edges: make(map[int][]Node),
	}

	for i := 1; i < v+1; i++ {
		graph.Nodes[i] = i
	}

	for i := 0; i < e; i++ {
		n1 := readStdin()
		n2 := readStdin()
		w := readStdin()

		graph.Edges[n1] = append(graph.Edges[n1], Node{n2, w})
		graph.Edges[n2] = append(graph.Edges[n2], Node{n1, w})
	}

	graph.prim()

	writer.WriteString(strconv.Itoa(totalWeight))
}

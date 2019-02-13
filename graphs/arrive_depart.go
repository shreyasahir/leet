package main

import "fmt"

type node struct {
	value int
}

type graph struct {
	nodes []node
	edges map[node][]node
}

func (g *graph) addNode(n node) {
	g.nodes = append(g.nodes, n)
}

func (n *node) string() string {
	return fmt.Sprintf("%v", n.value)
}
func (g *graph) addEdge(n1, n2 node) {
	if g.edges == nil {
		g.edges = make(map[node][]node)
	}
	g.edges[n1] = append(g.edges[n1], n2)
}

func showGraph(g *graph) {
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].string() + " -> "
		near := g.edges[g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].string() + " "
		}
		s += "\n"
	}
	fmt.Println("Graph is", s)
}

func fillGraph() {
	g := graph{}

	n0 := node{1}
	n1 := node{2}
	n2 := node{3}
	n3 := node{4}
	n4 := node{5}
	n5 := node{6}
	n6 := node{7}
	g.addNode(n0)
	g.addNode(n1)
	g.addNode(n3)
	g.addNode(n2)
	g.addNode(n4)
	g.addNode(n5)
	g.addNode(n6)

	g.addEdge(n0, n1)
	g.addEdge(n0, n2)
	g.addEdge(n1, n3)
	g.addEdge(n1, n4)
	g.addEdge(n2, n5)
	g.addEdge(n2, n6)

	showGraph(&g)
	g.arrivedepartdfs()
	// g.topologicalSort()
}

func main() {
	fillGraph()
}

type timings struct {
	arrival   int
	departure int
}

func (g *graph) arrivedepartdfs() {
	visited := make(map[node]bool, len(g.nodes))
	// timing := make(map[node]timings, len(g.nodes))
	arrival := make(map[node]int, len(g.nodes))
	departure := make(map[node]int, len(g.nodes))
	var stack []node
	stack = append(stack, g.nodes[0])
	count := 0
	for {
		if len(stack) == 0 {
			break
		}
		n := stack[len(stack)-1]
		arrival[n] = count
		count++
		stack = stack[:len(stack)-1]
		near := g.edges[n]
		for i := 0; i < len(near); i++ {
			j := near[i]
			if !visited[j] {
				stack = append(stack, j)
				visited[j] = true
			}
		}
		departure[n] = count
		count++
		fmt.Printf("%v\n", n)
	}

	for _, v := range g.nodes {
		fmt.Println("vertex", v)
		fmt.Println("arrival", arrival[v])
		fmt.Println("departure", departure[v])

	}
}

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

func (g *graph) dfs() {
	visited := make(map[node]bool, len(g.nodes))
	var stack []node
	stack = append(stack, g.nodes[0])

	for {
		if len(stack) == 0 {
			break
		}
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		near := g.edges[n]
		for i := 0; i < len(near); i++ {
			j := near[i]
			if !visited[j] {
				stack = append(stack, j)
				visited[j] = true
			}
		}
		fmt.Printf("%v\n", n)
	}
}

func (g *graph) topologicalSort() {

	var inDegree = make(map[node]int, len(g.nodes))
	for _, v := range g.nodes {
		inDegree[v] = 0
	}

	for _, v := range g.nodes {
		near := g.edges[v]
		for i := 0; i < len(near); i++ {
			j := near[i]
			inDegree[j]++
		}
	}

	var queue []node

	for v := range inDegree {
		if inDegree[v] == 0 {
			queue = append(queue, v)
		}
	}

	fmt.Println("queue with zero incoming edges", queue)
	count := 0
	var topSort []node

	for {
		if len(queue) == 0 {
			break
		}
		var u node
		u, queue = queue[0], queue[1:]
		topSort = append(topSort, u)
		near := g.edges[u]
		for i := 0; i < len(near); i++ {
			j := near[i]
			inDegree[j]--
			if inDegree[j] == 0 {
				queue = append(queue, j)
			}
		}
		count++
	}
	if count != len(g.nodes) {
		fmt.Println("There exist a cycle")
	} else {
		fmt.Println("Topological sort", topSort)
	}
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
	g.dfs()
	g.topologicalSort()
}

func main() {
	fillGraph()
}

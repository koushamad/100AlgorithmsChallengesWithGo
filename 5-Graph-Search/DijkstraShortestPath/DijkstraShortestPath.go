package DijkstraShortestPath

import DataStructure "github.com/koushamad/100AlgorithmsChallengesWithGo/0-Data-Structure"

type Graph struct {
	Vertexes []*Vertex
}

type Vertex struct {
	Key   string
	Edges map[string]*Vertex
	Costs map[string]int
	Paths []*Path
	Stack *DataStructure.StackLinkedList
}

type Path struct {
	key                        string
	Vertex                     *Vertex
	PreviousVertex             *Vertex
	ShortestDistanceFromVertex int
}

func NewGraph() *Graph {
	return &Graph{
		Vertexes: make([]*Vertex, 0),
	}
}

func NewVertex(key string) *Vertex {
	return &Vertex{
		Key:   key,
		Edges: make(map[string]*Vertex, 0),
		Paths: make([]*Path, 0),
		Costs: make(map[string]int),
		Stack: DataStructure.NewStackLinkedList(),
	}
}

func (pv *Vertex) NewPath(v *Vertex) *Path {
	return &Path{
		key:                        v.Key,
		Vertex:                     v,
		PreviousVertex:             nil,
		ShortestDistanceFromVertex: -1,
	}
}

func (g *Graph) AddVertex(key string) *Vertex {
	if g.Contain(key) {
		return g.GetVertex(key)
	}

	vertex := NewVertex(key)
	g.Vertexes = append(g.Vertexes, vertex)
	return vertex
}

func (g *Graph) AddEdge(from, to string, cost int) {
	fromVertex := g.AddVertex(from)
	toVertex := g.AddVertex(to)

	if fromVertex.Contain(to) {
		return
	}

	fromVertex.AddEdge(toVertex, cost)
}

func (g *Graph) AddEdge2(from, to string, cost int) {
	fromVertex := g.AddVertex(from)
	toVertex := g.AddVertex(to)

	if !fromVertex.Contain(to) {
		fromVertex.AddEdge(toVertex, cost)
	}

	if !toVertex.Contain(from) {
		toVertex.AddEdge(fromVertex, cost)
	}
}

func (g *Graph) Contain(key string) bool {
	for _, v := range g.Vertexes {
		if v.Key == key {
			return true
		}
	}
	return false
}

func (g *Graph) GetVertex(key string) *Vertex {
	for _, v := range g.Vertexes {
		if v.Key == key {
			return v
		}
	}
	return nil
}

func (pv *Vertex) Contain(key string) bool {
	for _, e := range pv.Edges {
		if e.Key == key {
			return true
		}
	}
	return false
}

func (pv *Vertex) AddEdge(edge *Vertex, cost int) {
	pv.Edges[edge.Key] = edge
	pv.Costs[edge.Key] = cost
}

func (pv *Vertex) GetPath(key string) *Path {
	for _, p := range pv.Paths {
		if p.key == key {
			return p
		}
	}
	return nil
}

func (pv *Vertex) AddPath(edge, pEdge *Vertex, cost int) {
	path := pv.GetPath(edge.Key)

	if path == nil {
		path = pv.NewPath(edge)
		pv.Paths = append(pv.Paths, path)
	}

	if path.PreviousVertex == nil || path.ShortestDistanceFromVertex > cost {
		path.ShortestDistanceFromVertex = cost
		path.PreviousVertex = pEdge
		pv.Costs[edge.Key] = cost
	}
}

func (pv *Vertex) CallPath(from, to *Vertex, cost int) int {
	if to.Key == from.Key {
		return cost
	}

	minCost := -1

	for key, edge := range from.Edges {

		if from.Stack.Contain(DataStructure.NewListItem(key)) || from.Key == key {
			continue
		}

		costEgge := from.Costs[key] + cost

		if pvCost, ok := pv.Costs[key]; ok {
			if pvCost < costEgge {
				continue
			}
		}

		pv.AddPath(edge, from, costEgge)
		from.Stack.Push(DataStructure.NewListItem(key))
		newCost := pv.CallPath(edge, to, costEgge)
		from.Stack.Pup()

		if newCost == -1 {
			continue
		}

		if minCost == -1 || newCost < minCost {
			minCost = newCost
			pv.AddPath(edge, from, newCost)
			from.Costs[key] = newCost
		}
	}

	return minCost
}

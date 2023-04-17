package ConnectedComponentsCount

type graph struct {
	vertexes []*vertex
}

type vertex struct {
	Key   int
	Edges map[int]*vertex
}

func NewVertex(key int) *vertex {
	return &vertex{
		Key:   key,
		Edges: make(map[int]*vertex, 0),
	}
}

func NewGraph() *graph {
	return &graph{
		vertexes: make([]*vertex, 0),
	}
}

func (g *graph) AddVertex(key int) *vertex {
	if !g.Contain(key) {
		vertex := NewVertex(key)
		g.vertexes = append(g.vertexes, vertex)
		return vertex
	}

	return g.GetVertex(key)
}

func (g *graph) AddEdge(from, to int) {
	fromVertex := g.AddVertex(from)
	toVertex := g.AddVertex(to)

	if fromVertex.Contain(to) {
		return
	}

	fromVertex.AddEdge(toVertex)
}

func (g *graph) AddEdge2(from, to int) {
	fromVertex := g.AddVertex(from)
	toVertex := g.AddVertex(to)

	if !fromVertex.Contain(to) {
		fromVertex.AddEdge(toVertex)
	}

	if toVertex.Contain(from) {
		toVertex.AddEdge(fromVertex)
	}
}

func (g *graph) Contain(key int) bool {
	for _, v := range g.vertexes {
		if v.Key == key {
			return true
		}
	}
	return false
}

func (g *graph) GetVertex(key int) *vertex {
	for _, v := range g.vertexes {
		if v.Key == key {
			return v
		}
	}
	return nil
}

func (v *vertex) Contain(key int) bool {
	for _, e := range v.Edges {
		if e.Key == key {
			return true
		}
	}
	return false
}

func (v *vertex) AddEdge(edge *vertex) {
	v.Edges[edge.Key] = edge
}

func solution(vertexes [][]int) int {
	gr := NewGraph()
	count := 0
	visits := make(map[int]bool, 0)

	for from, edges := range vertexes {
		for _, to := range edges {
			gr.AddEdge(from, to)
		}
	}

	lenVisits := len(visits)
	for _, ver := range gr.vertexes {
		visits = visit(ver, visits)
		if lenVisits < len(visits) {
			count++
		}
	}

	return count
}

func visit(vertex *vertex, visits map[int]bool) map[int]bool {

	if _, ok := visits[vertex.Key]; ok {
		return visits
	}

	visits[vertex.Key] = true

	for _, edge := range vertex.Edges {
		for _, ver := range edge.Edges {
			visits = visit(ver, visits)
		}
	}

	return visits
}

package DataStructure

type GraphMatrix struct {
	matrix [][]int
	keys   map[string]int
}

func NewGraphMatrix() *GraphMatrix {
	return &GraphMatrix{
		matrix: make([][]int, 0),
		keys:   make(map[string]int),
	}
}

func (g *GraphMatrix) AddVertex(key string) {
	if g.Contain(key) {
		return
	}

	capacity := len(g.matrix)
	g.matrix = append(g.matrix, make([]int, capacity))

	for i := 0; i <= capacity; i++ {
		g.matrix[i] = append(g.matrix[i], 0)
	}

	g.keys[key] = capacity
}

func (g *GraphMatrix) GetVertex(key string) []int {
	if !g.Contain(key) {
		return make([]int, 0)
	}

	index := g.keys[key]
	return g.matrix[index]
}

func (g *GraphMatrix) GetVertexKey(index int) string {

	for key, i := range g.keys {
		if i == index {
			return key
		}
	}

	return ""
}

func (g *GraphMatrix) GetVertexIndex(key string) int {
	if !g.Contain(key) {
		g.AddVertex(key)
	}

	return g.keys[key]
}

func (g *GraphMatrix) AddEdge(from, to string, cost int) {
	if !g.Contain(to) {
		g.AddVertex(to)
	}

	if !g.Contain(from) {
		g.AddVertex(from)
	}

	toIndex := g.GetVertexIndex(to)
	fromIndex := g.GetVertexIndex(from)

	g.matrix[toIndex][fromIndex] = cost
}

func (g *GraphMatrix) AddEdge2(from, to string, cost int) {
	if !g.Contain(to) {
		g.AddVertex(to)
	}

	if !g.Contain(from) {
		g.AddVertex(from)
	}

	toIndex := g.GetVertexIndex(to)
	fromIndex := g.GetVertexIndex(from)

	g.matrix[toIndex][fromIndex] = cost
	g.matrix[fromIndex][toIndex] = cost
}

func (g *GraphMatrix) GetEdgeCost(from, to string) int {
	if !g.Contain(to) {
		g.AddVertex(to)
	}

	if !g.Contain(from) {
		g.AddVertex(from)
	}

	toIndex := g.GetVertexIndex(to)
	fromIndex := g.GetVertexIndex(from)

	return g.matrix[toIndex][fromIndex]
}

func (g *GraphMatrix) Contain(key string) bool {
	if _, ok := g.keys[key]; ok {
		return true
	}

	return false
}

func (g *GraphMatrix) GetKeys() map[string]int {
	return g.keys
}

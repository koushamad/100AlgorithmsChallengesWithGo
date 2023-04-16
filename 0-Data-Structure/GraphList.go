package DataStructure

import (
	"fmt"
)

// GraphList Struct
type GraphList struct {
	vertices []*Vertex
}

// Vertex Struct
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// NewVertex create a new vertex
func NewVertex(key int) *Vertex {
	return &Vertex{key: key}
}

// NewGraphList create a new GraphList
func NewGraphList() *GraphList {
	return &GraphList{}
}

// InitGraphListByMatrix create the graph by matrix
func InitGraphListByMatrix(matrix [][]int) *GraphList {
	graph := NewGraphList()

	for keyVertex, vertices := range matrix {
		graph.AddVertex(keyVertex)
		for keyEdge, value := range vertices {
			if value > 0 {
				graph.AddEdge(keyVertex, keyEdge)
			}
		}
	}

	return graph
}

// InitGraphListByArray create the graph by array
func InitGraphListByArray(matrix [][]int) *GraphList {
	indexes := getKeyIndexesArray(matrix)
	graph := NewGraphList()
	lenX := len(matrix[0])
	lenY := len(matrix)

	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {

			key := getKeyArray(y, x)
			if _, ok := indexes[key]; !ok {
				continue
			}

			graph.AddVertex(key)

			//Check Left
			ketLeft := getKeyArray(y, x-1)
			if _, ok := indexes[ketLeft]; ok {
				graph.AddEdge2(key, ketLeft)
			}

			//Check Right
			ketRight := getKeyArray(y, x+1)
			if _, ok := indexes[ketRight]; ok {
				graph.AddEdge2(key, ketRight)
			}

			//Check Top
			keyTop := getKeyArray(y-1, x)
			if _, ok := indexes[keyTop]; ok {
				graph.AddEdge2(key, keyTop)
			}

			//Check Down
			keyDown := getKeyArray(y+1, x)
			if _, ok := indexes[keyDown]; ok {
				graph.AddEdge2(key, keyDown)
			}
		}
	}

	return graph
}

// AddVertex add a new vertex to graph
func (g *GraphList) AddVertex(key int) {
	if g.Contain(key) {
		return
	}

	g.vertices = append(g.vertices, NewVertex(key))
}

// AddEdge add new edge to graph
func (g *GraphList) AddEdge(from, to int) {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)

	if fromVertex == nil {
		fromVertex = NewVertex(from)
	}
	if toVertex == nil {
		toVertex = NewVertex(to)
	}

	fromVertex.AddVertex(toVertex)
}

// AddEdge2 add new 2 way edge to graph
func (g *GraphList) AddEdge2(from, to int) {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)

	if fromVertex == nil {
		fromVertex = NewVertex(from)
	}
	if toVertex == nil {
		toVertex = NewVertex(to)
	}

	fromVertex.AddVertex(toVertex)
	toVertex.AddVertex(fromVertex)
}

// Contain check key is exist in vertex
func (g *GraphList) Contain(key int) bool {
	for _, v := range g.vertices {
		if key == v.key {
			return true
		}
	}
	return false
}

// GetVertex find the vertex by key
func (g *GraphList) GetVertex(key int) *Vertex {
	for _, v := range g.vertices {
		if v.key == key {
			return v
		}
	}

	return nil
}

// GetEdge find the vertex by edges
func (g *GraphList) GetEdge(key int) []*Vertex {
	for _, v := range g.vertices {
		if v.key == key {
			return v.adjacent
		}
	}

	return []*Vertex{}
}

func (g *GraphList) Print() {
	for _, vertex := range g.vertices {
		fmt.Printf("vertex %d:", vertex.key)
		for _, edge := range vertex.adjacent {
			fmt.Printf(" %d", edge.key)
		}
		fmt.Println()
	}
}

// AddVertex add a new item to vertex
func (v *Vertex) AddVertex(ver *Vertex) {
	if v.Contain(ver.key) {
		return
	}

	v.adjacent = append(v.adjacent, ver)
}

// Contain check key is exist in vertex
func (v *Vertex) Contain(key int) bool {
	for _, vt := range v.adjacent {
		if vt.key == key {
			return true
		}
	}

	return false
}

// GetKey return the key of Vertex
func (v *Vertex) GetKey() int {
	return v.key
}

// GetAdjacent GetKey return the adjacent of Vertex
func (v *Vertex) GetAdjacent() []*Vertex {
	return v.adjacent
}

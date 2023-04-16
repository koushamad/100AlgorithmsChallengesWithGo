package DataStructure

import (
	"testing"
)

func TestInitGraphListByMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name          string
		args          args
		vertexContain [][]int
	}{
		{
			name: "Case 1",
			args: args{
				matrix: [][]int{
					0: {0: 0, 1: 0, 2: 0, 3: 1, 4: 0, 5: 1},
					1: {0: 0, 1: 0, 2: 1, 3: 0, 4: 0, 5: 0},
					2: {0: 1, 1: 1, 2: 0, 3: 0, 4: 0, 5: 1},
					3: {0: 1, 1: 1, 2: 1, 3: 1, 4: 1, 5: 0},
					4: {0: 1, 1: 1, 2: 1, 3: 0, 4: 0, 5: 0},
					5: {0: 0, 1: 1, 2: 0, 3: 1, 4: 0, 5: 0},
				},
			},
			vertexContain: [][]int{
				0: {3, 5},
				1: {2},
				2: {0, 1, 5},
				3: {0, 1, 2, 3, 4},
				4: {0, 1, 2},
				5: {3, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			graph := InitGraphListByMatrix(tt.args.matrix)
			graph.Print()

			for key, contains := range tt.vertexContain {
				vertex := graph.GetVertex(key)
				for _, edge := range contains {
					if got := vertex.Contain(edge); got != true {
						t.Errorf("GraphList Contain() = %v, want %v", got, true)
					}
				}
			}
		})
	}
}

func TestInitGraphListByArray(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name          string
		args          args
		vertexContain [][]int
	}{
		{
			name: "Case 1",
			args: args{
				matrix: [][]int{
					0: {0: 2, 1: 0, 2: 1, 3: 1, 4: 1, 5: 1},
					1: {0: 1, 1: 0, 2: 1, 3: 0, 4: 0, 5: 1},
					2: {0: 1, 1: 0, 2: 1, 3: 0, 4: 1, 5: 1},
					3: {0: 1, 1: 0, 2: 1, 3: 0, 4: 1, 5: 1},
					4: {0: 1, 1: 0, 2: 1, 3: 0, 4: 1, 5: 1},
					5: {0: 1, 1: 1, 2: 1, 3: 0, 4: 3, 5: 1},
				},
			},
			vertexContain: [][]int{
				0:  {10},
				2:  {3, 12},
				45: {44, 35, 55},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			graph := InitGraphListByArray(tt.args.matrix)
			graph.Print()

			for key, contains := range tt.vertexContain {
				vertex := graph.GetVertex(key)
				for _, edge := range contains {
					if got := vertex.Contain(edge); got != true {
						t.Errorf("GraphList Contain() = %v, want %v", got, true)
					}
				}
			}
		})
	}
}

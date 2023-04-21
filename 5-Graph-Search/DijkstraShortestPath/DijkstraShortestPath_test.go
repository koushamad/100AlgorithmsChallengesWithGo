package DijkstraShortestPath

import (
	"reflect"
	"testing"
)

func TestVertex_CallPath(t *testing.T) {
	type fields struct {
		graph *Graph
	}
	type args struct {
		previousCost int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]int
	}{
		{
			name: "Case 1",
			fields: fields{
				graph: func() *Graph {
					graph := NewGraph()
					graph.AddEdge2("A", "B", 6)
					graph.AddEdge2("A", "D", 1)
					graph.AddEdge2("B", "E", 2)
					graph.AddEdge2("B", "C", 5)
					graph.AddEdge2("B", "D", 2)
					graph.AddEdge2("D", "E", 1)
					graph.AddEdge2("E", "C", 5)
					return graph
				}(),
			},
			args: args{
				previousCost: 0,
			},
			want: map[string]int{
				"A": 1,
				"B": 1,
				"C": 1,
				"D": 1,
				"E": 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pv := tt.fields.graph.GetVertex("A")
			if got := pv.CallPath(pv, tt.fields.graph.GetVertex("C"), tt.args.previousCost); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CallPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

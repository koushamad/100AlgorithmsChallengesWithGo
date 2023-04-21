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
		from         string
		to           string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
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
				from:         "A",
				to:           "C",
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pv := tt.fields.graph.GetVertex(tt.args.from)
			if got := pv.CallPath(pv, tt.fields.graph.GetVertex(tt.args.to), tt.args.previousCost); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CallPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

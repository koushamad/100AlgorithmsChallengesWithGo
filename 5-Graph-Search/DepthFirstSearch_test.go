package GraphSearch

import (
	DataStructure "github.com/koushamad/100AlgorithmsChallengesWithGo/0-Data-Structure"
	"testing"
)

func TestDepthFirstSearch_HashPath(t *testing.T) {
	type fields struct {
		graph *DataStructure.GraphList
		stack *DataStructure.StackLinkedList
	}
	type args struct {
		src int
		dst int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Case 1",
			fields: fields{
				graph: func() *DataStructure.GraphList {
					graph := DataStructure.NewGraphList()
					graph.AddVertex(1)
					graph.AddVertex(2)
					graph.AddVertex(3)
					graph.AddVertex(4)
					graph.AddVertex(5)
					graph.AddVertex(6)
					graph.AddEdge(1, 2)
					graph.AddEdge(1, 4)
					graph.AddEdge(2, 3)
					graph.AddEdge(4, 2)
					graph.AddEdge(4, 6)
					graph.AddEdge(5, 4)
					return graph
				}(),
				stack: DataStructure.NewStackLinkedList(),
			},
			args: args{
				src: 1,
				dst: 6,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DepthFirstSearch{
				graph: tt.fields.graph,
				stack: tt.fields.stack,
			}
			if got := d.HashPath(tt.fields.graph.GetVertex(tt.args.src), tt.fields.graph.GetVertex(tt.args.dst)); got != tt.want {
				t.Errorf("HashPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

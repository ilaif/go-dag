package dag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectedComponentRoots(t *testing.T) {
	g := &Graph{
		Nodes: []Node{0, 1, 2, 3},
		Edges: []Edge{
			{Depender: 2, Dependee: 0},
			{Depender: 3, Dependee: 1},
		},
	}
	assert.ElementsMatch(t, g.ConnectedComponentRoots(), []Node{0, 1})

	g = &Graph{
		Nodes: []Node{0, 1, 2, 3, 4, 5, 6},
		Edges: []Edge{
			{Depender: 2, Dependee: 0},
			{Depender: 3, Dependee: 1},
			{Depender: 4, Dependee: 2},
			{Depender: 5, Dependee: 2},
			{Depender: 6, Dependee: 5},
		},
	}
	assert.ElementsMatch(t, g.ConnectedComponentRoots(), []Node{0, 1})
}

func TestDirectDependers(t *testing.T) {
	g := &Graph{
		Nodes: []Node{0, 1, 2, 3, 4, 5, 6},
		Edges: []Edge{
			{Depender: 2, Dependee: 0},
			{Depender: 3, Dependee: 1},
			{Depender: 4, Dependee: 2},
			{Depender: 5, Dependee: 2},
			{Depender: 6, Dependee: 5},
		},
	}
	assert.ElementsMatch(t, g.DirectDependers(2), []Node{4, 5})
}

func TestDirectDependees(t *testing.T) {
	g := &Graph{
		Nodes: []Node{0, 1, 2, 3, 4, 5, 6},
		Edges: []Edge{
			{Depender: 2, Dependee: 0},
			{Depender: 3, Dependee: 1},
			{Depender: 4, Dependee: 2},
			{Depender: 5, Dependee: 2},
			{Depender: 6, Dependee: 5},
		},
	}
	assert.ElementsMatch(t, g.DirectDependees(2), []Node{0})
}

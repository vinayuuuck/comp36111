package structures

import (
    "fmt"
)

// Directed graph will be represented by an adjacency list
type Graph struct {
    Vertices []*Vertex
}

// Vertex will be represented by a string ID and a list of adjacent vertices, that is, edges directed from this vertex to other vertices
type Vertex struct {
    ID int
    Adjacent []*Vertex
}

// AddEdge adds an edge to the graph, that is, adds a given vertex to the adjacency list of another vertex
func (g *Graph) AddEdge(v1, v2 int) {
    vertex1 := g.GetVertex(v1)
    vertex2 := g.GetVertex(v2)
    vertex1.Adjacent = append(vertex1.Adjacent, vertex2)
}

// AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(v *Vertex) {
    g.Vertices = append(g.Vertices, v)
}

// GetVertex returns a vertex with a given ID
func (g *Graph) GetVertex(id int) *Vertex {
    for _, v := range g.Vertices {
        if v.ID == id {
            return v
        }
    }
    return nil
}

// GetVertices returns all connected vertices to a given vertex
func (g *Graph) GetVertices(v *Vertex) []*Vertex {
    return v.Adjacent
}

// Print prints the graph
func (g *Graph) Print() {
    for _, v := range g.Vertices {
        fmt.Printf("%d: ", v.ID)
        for _, v2 := range v.Adjacent {
            fmt.Printf("%d ", v2.ID)
        }
        fmt.Printf("\n")
    }
}

// Constructor for a graph. Given an integer n, it creates a graph with n vertices, with no connections between them
func NewGraph(n int) *Graph {
    var g Graph
    for i := 0; i < n; i++ {
        g.AddVertex(&Vertex{ID: i})
    }
    return &g
}

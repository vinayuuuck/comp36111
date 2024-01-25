package structures

import (
    "fmt"
)

// Directed graph will be represented by an adjacency list
type Graph struct {
    Vertices []*Vertex
    Edges []*Edge
}

// Vertex will be represented by an int ID.
type Vertex struct {
    ID int
}

// Edge will be a hash table with a key of two vertices and a value of the capacity of the edge, and flow through the edge which will be initialized to 0 by default
type Edge struct {
    Vertices [2]*Vertex
    Capacity int
    Flow int
}

// The graph will be initialized with a list of vertices. This will just mean that the graph contains n vertices, but no edges yet
func NewGraph(n int) *Graph{
    var g Graph
    for i := 0; i < n; i++ {
        g.Vertices = append(g.Vertices, &Vertex{ID: i})
    }
    return &g
}

// Add an edge to the graph. The arguments will be v1 int, v2 int, and capacity int
func (g *Graph) AddEdge(v1 int, v2 int, capacity int) {
    var e Edge
    e.Vertices[0] = g.Vertices[v1]
    e.Vertices[1] = g.Vertices[v2]
    e.Capacity = capacity
    g.Edges = append(g.Edges, &e)
}

// Print the graph
func (g *Graph) Print() {
    fmt.Println("Vertices:")
    for _, v := range g.Vertices {
        fmt.Println(v.ID)
    }
    fmt.Println("Edges:")
    for _, e := range g.Edges {
        fmt.Println("Edge from", e.Vertices[0].ID, "to", e.Vertices[1].ID, "with capacity", e.Capacity)
    }
}

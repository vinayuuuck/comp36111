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
    ID string
    Adjacent []*Vertex
}

// AddEdge adds an edge to the graph, that is, adds a given vertex to the adjacency list of another vertex
func (g *Graph) AddEdge(v1, v2 *Vertex) {
    v1.Adjacent = append(v1.Adjacent, v2)
}

// if addEdge is called with integers, it will find the vertices with the given IDs and call the function above
func (g *Graph) AddEdgeInt(v1, v2 int) {
    g.AddEdge(g.GetVertex(fmt.Sprintf("%d", v1)), g.GetVertex(fmt.Sprintf("%d", v2)))
}

// AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(v *Vertex) {
    g.Vertices = append(g.Vertices, v)
}

// GetVertex returns a vertex with a given ID
func (g *Graph) GetVertex(id string) *Vertex {
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
        fmt.Printf("%s -> ", v.ID)
        for _, v := range v.Adjacent {
            fmt.Printf("%s ", v.ID)
        }
        fmt.Println()
    }
}

// Constructor for a graph. Given an integer n, it creates a graph with n vertices, with no connections between them
func NewGraph(n int) *Graph {
    var g Graph
    for i := 0; i < n; i++ {
        var v Vertex
        v.ID = fmt.Sprintf("%d", i)
        g.AddVertex(&v)
    }
    return &g
}

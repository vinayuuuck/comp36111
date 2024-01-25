package main

import (
    "internal/structures"
    "internal/algorithms"
)

func main() {
    // Create a graph
    // First define the number of vertices
    var n int = 5
    // Then create a graph with n vertices
    g := structures.NewGraph(n)
    // Add edges
    g.AddEdge(0, 1)
    g.AddEdge(0, 2)
    g.AddEdge(1, 2)
    // Print the graph
    g.Print()
    algorithms.FordFulkerson(g)
}

package main

import (
    "fmt"
    // "github.com/vinayuuuck/comp36111/algorithms"
    "github.com/vinayuuuck/comp36111/structures"
)

func main() {
    fmt.Println("Hello World!")
    // Create a graph
    // First define the number of vertices
    var n int = 5
    // Then create a graph with n vertices
    g := structures.NewGraph(n)
    // Add edges
    g.AddEdgeInt(0, 1)
    g.AddEdgeInt(0, 2)
    g.AddEdgeInt(0, 3)
    // Print the graph
    g.Print()
}

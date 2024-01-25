package main

import (
    "internal/structures"
    "internal/algorithms"
)

func main() {

    // Test the Ford-Fulkerson algorithm
    // We will assume that v0 will always be the source and vn-1 will always be the sink
    var n int = 5
    g := structures.NewGraph(n)
    g.AddEdge(0, 1, 10)
    g.AddEdge(0, 2, 12)
    g.AddEdge(1, 2, 9)
    g.AddEdge(1, 3, 21)
    g.AddEdge(2, 3, 16)
    g.AddEdge(2, 4, 88)
    g.AddEdge(3, 4, 15)

    g.Print()
    algorithms.FlowMax(g)
}

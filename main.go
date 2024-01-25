package main

import (
    "internal/structures"
    "internal/algorithms"
)

func main() {

    // Test the Ford-Fulkerson algorithm
    var n int = 5
    g := structures.NewGraph(n)
    g.AddEdge(0, 1, 1)
    g.Print()
    algorithms.FordFulkerson(g)
}

package main

import (
    "fmt"
    "internal/structures"
    "internal/algorithms"
)

func main() {

    // Test the Ford-Fulkerson algorithm
    // We will assume that v0 will always be the source and vn-1 will always be the sink
    var n int = 5
    s := 0
    t := n - 1
    g := structures.NewGraph(n)
    g.AddEdge(s, 1, 10)
    g.AddEdge(0, 2, 12)
    g.AddEdge(1, 2, 9)
    g.AddEdge(1, 3, 21)
    g.AddEdge(2, 3, 16)
    g.AddEdge(2, t, 88)
    g.AddEdge(3, t, 15)

    g.Print()
    flow := algorithms.FlowMax(g)
    fmt.Println(flow)
}

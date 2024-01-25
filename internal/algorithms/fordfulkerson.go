package algorithms

import (
    "fmt"
    "internal/structures"
)

// Take a graph as input and call the print function
func FordFulkerson(graph *structures.Graph) {
    fmt.Println("Ford Fulkerson")
    graph.Print()
}

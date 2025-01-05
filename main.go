 package main

import (
    "fmt"
)

func main() {
    // Create nodes
    nodes := []*Node{
        NewNode(1, "localhost:8001"),
        NewNode(2, "localhost:8002"),
        NewNode(3, "localhost:8003"),
    }

    // Connect nodes
    nodes[0].Join(nodes[1])
    nodes[1].Join(nodes[2])

    // Start visualization
    StartVisualization(nodes)

    // Start listening on each node
    for _, node := range nodes {
        go node.Listen()
    }

    // Demonstrate some operations
    fmt.Println("Storing data across nodes...")
    nodes[0].Store("key1", "value1")
    nodes[1].Store("key2", "value2")
    nodes[2].Store("key3", "value3")

    // Keep the program running
    select {}
}
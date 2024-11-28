package main

import "fmt"

func main() {
	// Create nodes
	node1 := NewNode(1, "localhost:8000")
	node2 := NewNode(2, "localhost:8001")
	node3 := NewNode(3, "localhost:8002")

	// Join node2 to node1
	node2.Join(node1)

	// Join node3 to node2
	node3.Join(node2)

	// Print neighbors for each node
	fmt.Println("Node 1 Neighbors:", node1.Neighbors)
	fmt.Println("Node 2 Neighbors:", node2.Neighbors)
	fmt.Println("Node 3 Neighbors:", node3.Neighbors)
}

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
	fmt.Println("Node 1 Neighbors:")
	for _, neighbor := range node1.Neighbors {
		// Dereference the pointer and print ID and Address
		fmt.Printf("{ID: %d Address: %s} ", neighbor.ID, neighbor.Address)
	}
	fmt.Println()

	fmt.Println("Node 2 Neighbors:")
	for _, neighbor := range node2.Neighbors {
		// Dereference the pointer and print ID and Address
		fmt.Printf("{ID: %d Address: %s} ", neighbor.ID, neighbor.Address)
	}
	fmt.Println()

	fmt.Println("Node 3 Neighbors:")
	for _, neighbor := range node3.Neighbors {
		// Dereference the pointer and print ID and Address
		fmt.Printf("{ID: %d Address: %s} ", neighbor.ID, neighbor.Address)
	}
	fmt.Println()
}

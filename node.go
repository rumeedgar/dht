package main

import "fmt"

// Node struct definition
type Node struct {
	ID        uint32
	Address   string
	Data      map[string]string
	Neighbors []Node
}

// NewNode creates a new Node instance
func NewNode(id uint32, address string) *Node {
	return &Node{
		ID:        id,
		Address:   address,
		Data:      make(map[string]string),
		Neighbors: []Node{},
	}
}

// Join method to add a node to the network
func (n *Node) Join(existingNode *Node) {
	// Find closest neighbors from the existing node
	closestNode := existingNode.FindClosest(n.ID)
	if closestNode == nil {
		// Handle the case where no closest node is found
		fmt.Println("No closest node found, adding this node as a neighbor.")
		n.Neighbors = append(n.Neighbors, *existingNode)
	} else {
		// Otherwise, use the closest node found
		n.Neighbors = append(n.Neighbors, *closestNode)
	}

	// Add the joining node to the existing node's neighbor list
	existingNode.Neighbors = append(existingNode.Neighbors, *n)

	// Notify the neighbors about the new node
	for _, neighbor := range n.Neighbors {
		neighbor.Neighbors = append(neighbor.Neighbors, *n)
	}
}

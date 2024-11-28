package main

import "fmt"

// Node represents a node in the network
type Node struct {
	ID        uint32
	Address   string
	Data      map[string]string
	Neighbors []*Node // Use pointers to Node for easier referencing
}

// NewNode creates a new node with the specified ID and address
func NewNode(id uint32, address string) *Node {
	return &Node{
		ID:        id,
		Address:   address,
		Data:      make(map[string]string),
		Neighbors: []*Node{}, // Empty slice of pointers
	}
}

func (n *Node) Join(existingNode *Node) error {
	// Check if the existing node is valid
	if existingNode == nil {
		return fmt.Errorf("existing node is nil")
	}

	// Find the closest node using FindClosest (now moved to routing.go)
	closestNode := FindClosest(existingNode, n.ID)
	if closestNode == nil {
		// Handle the case where no closest node is found (first node in the network)
		fmt.Println("No closest node found, adding this node as a neighbor.")
		// Add only if not already in the neighbors list
		if !contains(existingNode.Neighbors, n) {
			existingNode.Neighbors = append(existingNode.Neighbors, n)
		}
	} else {
		// Otherwise, add the closest node found
		fmt.Printf("Adding closest node (ID: %d) as a neighbor\n", closestNode.ID)
		// Add only if not already in the neighbors list
		if !contains(existingNode.Neighbors, closestNode) {
			existingNode.Neighbors = append(existingNode.Neighbors, closestNode)
		}
	}

	// Add the joining node to the existing node's neighbor list
	if !contains(n.Neighbors, existingNode) {
		n.Neighbors = append(n.Neighbors, existingNode)
	}

	// Notify the neighbors about the new node
	for _, neighbor := range n.Neighbors {
		// Make sure the current node is added to the neighbor's neighbor list
		if !contains(neighbor.Neighbors, n) {
			neighbor.Neighbors = append(neighbor.Neighbors, n)
		}
	}

	// Debug print to show the neighbors after joining
	fmt.Printf("Node %d Neighbors: %+v\n", n.ID, n.Neighbors)

	return nil
}

// Helper function to check if a slice contains a node
func contains(neighbors []*Node, node *Node) bool {
	for _, neighbor := range neighbors {
		if neighbor.ID == node.ID {
			return true
		}
	}
	return false
}

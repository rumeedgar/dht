package main

import "sort"

// XOR distance between two node IDs
func xorDistance(a, b uint32) uint32 {
	return a ^ b
}

// KBucket represents a list of nodes sorted by XOR distance
type KBucket struct {
	Nodes []*Node // Use pointers to Node for easier referencing
}

// AddNode to KBucket, maintaining the sorted order based on XOR distance
func (kb *KBucket) AddNode(newNode *Node, targetID uint32) {
	// Sort nodes by XOR distance (ascending)
	sort.Slice(kb.Nodes, func(i, j int) bool {
		return xorDistance(kb.Nodes[i].ID, targetID) < xorDistance(kb.Nodes[j].ID, targetID)
	})
	kb.Nodes = append(kb.Nodes, newNode)
}

// FindClosest finds the closest node based on XOR distance (now in routing.go)
func FindClosest(n *Node, targetID uint32) *Node {
	var closest *Node
	minDistance := uint32(^uint32(0)) // Maximum possible uint32 value

	// Loop through neighbors to find the closest
	for _, neighbor := range n.Neighbors {
		distance := xorDistance(neighbor.ID, targetID)
		if distance < minDistance {
			minDistance = distance
			closest = neighbor
		}
	}

	// Return the closest node or nil if no neighbors exist
	return closest
}

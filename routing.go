package main

import (
	"fmt"
	"sort"
)

// XOR distance between two node IDs
func xorDistance(a, b uint32) uint32 {
	return a ^ b
}

// KBucket represents a list of nodes sorted by XOR distance
type KBucket struct {
	Nodes    []*Node
	Capacity int
}

// NewKBucket creates a new KBucket with a specified capacity
func NewKBucket(capacity int) *KBucket {
	return &KBucket{
		Nodes:    []*Node{},
		Capacity: capacity,
	}
}

// AddNode adds a node to the KBucket, maintaining the sorted order based on XOR distance
func (kb *KBucket) AddNode(newNode *Node, targetID uint32) {
	if len(kb.Nodes) < kb.Capacity {
		kb.Nodes = append(kb.Nodes, newNode)
		return
	}

	sort.Slice(kb.Nodes, func(i, j int) bool {
		return xorDistance(kb.Nodes[i].ID, targetID) < xorDistance(kb.Nodes[j].ID, targetID)
	})

	if xorDistance(newNode.ID, targetID) < xorDistance(kb.Nodes[len(kb.Nodes)-1].ID, targetID) {
		kb.Nodes[len(kb.Nodes)-1] = newNode
	}
}

// FindClosestKNodes returns the closest K nodes based on XOR distance
func (kb *KBucket) FindClosestKNodes(targetID uint32, k int) []*Node {
	sort.Slice(kb.Nodes, func(i, j int) bool {
		return xorDistance(kb.Nodes[i].ID, targetID) < xorDistance(kb.Nodes[j].ID, targetID)
	})

	if len(kb.Nodes) < k {
		return kb.Nodes
	}
	return kb.Nodes[:k]
}

// Print displays the contents of the KBucket
func (kb *KBucket) Print() {
	for _, node := range kb.Nodes {
		fmt.Printf("Node %d at Address %s\n", node.ID, node.Address)
	}
}

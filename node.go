package main

import (
	"sync"
)

// Node represents a node in the distributed hash table network
type Node struct {
	ID           uint32
	Address      string
	Neighbors    []*Node
	RoutingTable *KBucket
	Data         map[string]string
	mu           sync.RWMutex
}

// NewNode creates a new node with specified ID and address
func NewNode(id uint32, address string) *Node {
	return &Node{
		ID:           id,
		Address:      address,
		Neighbors:    make([]*Node, 0),
		RoutingTable: NewKBucket(8),
		Data:         make(map[string]string),
	}
}

// Store adds a key-value pair to the node's data
func (n *Node) Store(key, value string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.Data[key] = value
}

// Retrieve fetches a value for a given key
func (n *Node) Retrieve(key string) (string, bool) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	value, found := n.Data[key]
	return value, found
}

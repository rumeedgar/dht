package main

type Node struct {
	ID        uint32
	Address   string
	Data      map[string]string // Key-value storage
	Neighbors []Node            // List of neighboring nodes
}

func NewNode(id uint32, address string) *Node {
	return &Node{
		ID:        id,
		Address:   address,
		Data:      make(map[string]string),
		Neighbors: []Node{},
	}
}

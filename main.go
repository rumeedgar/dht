package main

import (
	"fmt"
	"time"
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

	// Start listening on each node
	for _, node := range nodes {
		go node.Listen()
	}

	// Demonstrate some operations
	fmt.Println("Storing data across nodes...")
	nodes[0].Store("key1", "value1")
	nodes[1].Store("key2", "value2")
	nodes[2].Store("key3", "value3")

	// Retrieve data
	fmt.Println("Retrieving data...")
	for _, node := range nodes {
		for _, key := range []string{"key1", "key2", "key3"} {
			value, found := node.Retrieve(key)
			if found {
				fmt.Printf("Node %d: %s = %s\n", node.ID, key, value)
			}
		}
	}

	// Keep the program running and show node status
	fmt.Println("Demonstrating network status...")
	for _, node := range nodes {
		fmt.Printf("Node %d Neighbors: %d\n", node.ID, len(node.Neighbors))
	}

	// Wait a bit to allow network operations
	time.Sleep(5 * time.Second)
}

package main

import (
	"fmt"
	"strings"
	"time"
)

// VisualizeNetwork collects the current state of the DHT network
func (n *Node) VisualizeNetwork(builder *strings.Builder) {
	builder.WriteString(fmt.Sprintf("\nNode %d (%s)\n", n.ID, n.Address))
	builder.WriteString(strings.Repeat("-", 50) + "\n")

	// Print routing table
	builder.WriteString("\n📋 Routing Table:\n")
	if len(n.Neighbors) == 0 {
		builder.WriteString("  └─ No neighbors\n")
	} else {
		for _, neighbor := range n.Neighbors {
			distance := xorDistance(n.ID, neighbor.ID)
			builder.WriteString(fmt.Sprintf("  └─ Node %d at %s (XOR Distance: %d)\n", neighbor.ID, neighbor.Address, distance))
		}
	}

	// Print k-bucket information
	builder.WriteString(fmt.Sprintf("\n🪣 K-Bucket (Capacity: %d)\n", n.RoutingTable.Capacity))
	if len(n.RoutingTable.Nodes) == 0 {
		builder.WriteString("  └─ Empty bucket\n")
	} else {
		for _, node := range n.RoutingTable.Nodes {
			builder.WriteString(fmt.Sprintf("  └─ Node %d\n", node.ID))
		}
	}

	// Print stored data
	builder.WriteString("\n💾 Stored Data:\n")
	if len(n.Data) == 0 {
		builder.WriteString("  └─ No data stored\n")
	} else {
		for key, value := range n.Data {
			builder.WriteString(fmt.Sprintf("  └─ %s: %s\n", key, value))
		}
	}

	builder.WriteString(strings.Repeat("-", 50) + "\n")
}

// StartVisualization periodically updates network visualization
func StartVisualization(nodes []*Node) {
	ticker := time.NewTicker(2 * time.Second)
	go func() {
		for range ticker.C {
			var visualization strings.Builder
			visualization.WriteString("\033[H\033[2J") // Clear screen
			visualization.WriteString(networkBanner()) // Collect banner as a string

			// Display all nodes
			for _, node := range nodes {
				node.VisualizeNetwork(&visualization)
				visualization.WriteString("\n")
			}

			// Network statistics
			visualization.WriteString("\n📊 Network Statistics:\n")
			visualization.WriteString(fmt.Sprintf("Total Nodes: %d\n", len(nodes)))
			totalData := 0
			for _, node := range nodes {
				totalData += len(node.Data)
			}
			visualization.WriteString(fmt.Sprintf("Total Stored Items: %d\n", totalData))

			// Print the complete visualization
			fmt.Print(visualization.String())
		}
	}()
}

// Returns the network banner as a string
func networkBanner() string {
	return `
+------------------------------------------+
|             DHT Network Status           |
+------------------------------------------+
`
}

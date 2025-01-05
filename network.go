package main

import (
    "encoding/json"
    "fmt"
    "net"
)

// Join adds a node as a neighbor and updates routing tables
func (n *Node) Join(otherNode *Node) {
    n.Neighbors = append(n.Neighbors, otherNode)
    otherNode.Neighbors = append(otherNode.Neighbors, n)

    n.RoutingTable.AddNode(otherNode, n.ID)
    otherNode.RoutingTable.AddNode(n, otherNode.ID)

    fmt.Printf("Node %d joined Node %d\n", n.ID, otherNode.ID)
}

// Listen handles incoming network connections
func (n *Node) Listen() {
    listener, err := net.Listen("tcp", n.Address)
    if err != nil {
        fmt.Printf("Error listening: %v\n", err)
        return
    }
    defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Printf("Error accepting connection: %v\n", err)
            continue
        }
        go n.handleConnection(conn)
    }
}

func (n *Node) StoreWithHash(key, value string) {
    hashedKey := hashKey(key)
    // Use hashedKey for routing or storage logic
    n.Store(hashedKey.String(), value)
}

// handleConnection processes incoming network requests
func (n *Node) handleConnection(conn net.Conn) {
    defer conn.Close()

    var request map[string]string
    decoder := json.NewDecoder(conn)
    if err := decoder.Decode(&request); err != nil {
        fmt.Printf("Error decoding request: %v\n", err)
        return
    }

    switch request["action"] {
    case "store":
        n.Store(request["key"], request["value"])
        fmt.Fprintf(conn, "Stored key: %s", request["key"])
    case "retrieve":
        value, found := n.Retrieve(request["key"])
        if found {
            fmt.Fprintf(conn, "Value: %s", value)
        } else {
            fmt.Fprintf(conn, "Key not found")
        }
    }
}

// FindClosest locates the nearest node to a target ID
func (n *Node) FindClosest(targetID uint32) *Node {
    var closest *Node
    minDistance := uint32(^uint32(0))

    for _, neighbor := range n.Neighbors {
        distance := xorDistance(neighbor.ID, targetID)
        if distance < minDistance {
            minDistance = distance
            closest = neighbor
        }
    }

    return closest
}
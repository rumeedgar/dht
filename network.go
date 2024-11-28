package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func (n *Node) Listen() {
	listener, _ := net.Listen("tcp", n.Address)
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go n.handleConnection(conn)
	}
}

func (n *Node) handleConnection(conn net.Conn) {
	defer conn.Close()
	decoder := json.NewDecoder(conn)
	var request map[string]string
	decoder.Decode(&request)

	if request["action"] == "store" {
		n.Store(request["key"], request["value"])
		fmt.Fprintf(conn, "Stored key: %s", request["key"])
	} else if request["action"] == "retrieve" {
		value, found := n.Retrieve(request["key"])
		if found {
			fmt.Fprintf(conn, "Value: %s", value)
		} else {
			fmt.Fprintf(conn, "Key not found")
		}
	}
}

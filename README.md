# Distributed Hash Table (DHT) Simulator

A Golang implementation of a simplified DHT network simulator that demonstrates the core concepts of distributed hash tables, similar to those used in peer-to-peer networks like BitTorrent and IPFS.

## Features

- 🌐 Node Network Simulation
  - Dynamic node joining and routing
  - XOR-based distance metrics for node organization
  - K-bucket routing tables
  - TCP-based inter-node communication

- 💾 Data Management
  - Distributed key-value storage
  - SHA-1 based key hashing
  - Data retrieval across nodes

- 📊 Real-time Visualization
  - Live network topology display
  - Node connection status
  - Routing table information
  - Stored data visualization

## Getting Started

### Prerequisites

- Go 1.15 or higher

### Installation

1. Clone the repository
```bash
git clone https://github.com/yourusername/dht-simulator.git
cd dht-simulator
```

2. Run the simulator
```bash
go run .
```

## Project Structure

- `main.go` - Entry point and network initialization
- `node.go` - Node implementation and data storage
- `routing.go` - DHT routing logic and k-bucket implementation
- `network.go` - Network communication and node joining
- `hash.go` - Key hashing functionality
- `visualization.go` - Real-time network visualization

## How It Works

The simulator creates a network of nodes where:
1. Each node has a unique ID and network address
2. Nodes are organized using XOR distance metrics
3. Routing tables (k-buckets) maintain network topology
4. Data is stored and retrieved across the network
5. Live visualization shows network status and data distribution

Example of node initialization and data storage:
```go
// Create nodes
nodes := []*Node{
    NewNode(1, "localhost:8001"),
    NewNode(2, "localhost:8002"),
    NewNode(3, "localhost:8003"),
}

// Connect nodes
nodes[0].Join(nodes[1])
nodes[1].Join(nodes[2])

// Store data
nodes[0].Store("key1", "value1")
```

## Technical Details

### Distance Calculation
The simulator uses XOR distance metrics to determine node proximity:
```go
func xorDistance(a, b uint32) uint32 {
    return a ^ b
}
```

### K-Bucket Routing
Implements a simplified version of Kademlia's k-bucket routing table:
- Maintains lists of known nodes
- Organizes nodes by XOR distance
- Facilitates efficient node lookup

### Data Storage
Uses a distributed key-value store with:
- SHA-1 based key hashing
- Distributed data storage across nodes
- Local caching for improved performance

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Future Improvements

- [ ] Implement node departure handling
- [ ] Add data replication across nodes
- [ ] Improve fault tolerance
- [ ] Add network partition handling
- [ ] Implement concurrent data operations
- [ ] Add more sophisticated routing algorithms

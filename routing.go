package main

func xorDistance(a, b uint32) uint32 {
    return a ^ b
}

func (n *Node) FindClosest(targetID uint32) *Node {
    var closest *Node
    minDistance := uint32(^uint32(0)) // Max uint32 value

    for _, neighbor := range n.Neighbors {
        distance := xorDistance(neighbor.ID, targetID)
        if distance < minDistance {
            minDistance = distance
            closest = &neighbor
        }
    }
    return closest
}

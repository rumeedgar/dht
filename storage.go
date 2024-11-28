package main

func (n *Node) Store(key, value string) {
	n.Data[key] = value
}

func (n *Node) Retrieve(key string) (string, bool) {
	val, exists := n.Data[key]
	return val, exists
}

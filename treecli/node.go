package treecli

type Node struct {
	Name   string
	Weight int
	Left   *Node
	Right  *Node
}

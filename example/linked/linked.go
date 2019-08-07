package main

type LinkNode struct {
	p    *LinkNode
	data interface{}
}
type ListLinkNode struct {
	head *LinkNode
	len  uint
}

func NewLinkNode(p *LinkNode, value int) *LinkNode {
	return &LinkNode{p, value}
}

func NewListLinkNode() ListLinkNode {
	return ListLinkNode{&LinkNode{nil, nil}, 0}
}

func (l ListLinkNode) Add(l *LinkNode, value int) {

}
func main() {

}

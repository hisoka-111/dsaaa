package list

type Node struct {
	Element interface{}
	Next    *Node
}

type List interface {
	IsEmpty() bool
	IsLast(*Node) bool
	Find(interface{}) *Node
	Delete(interface{})
	FindPrevious(interface{}) *Node
	Insert(interface{}, *Node)
	DeleteList()
	Header() *Node
	First() *Node
	Advance(*Node) *Node
	Retrieve(*Node) interface{}
}

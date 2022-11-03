package linkedlist

type LinkedList interface {
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

type Node struct {
	Element interface{}
	Next    *Node
}

type MyList struct {
	head *Node
}

func NewLinkedList() LinkedList {
	return &MyList{
		head: &Node{},
	}
}

func (l *MyList) IsEmpty() bool {
	return l.head.Next == nil
}

func (l *MyList) IsLast(position *Node) bool {
	if position == nil {
		return false
	}
	return position.Next == nil
}

func (l *MyList) Find(x interface{}) *Node {
	for cur := l.head.Next; cur != nil; cur = cur.Next {
		if cur.Element == x {
			return cur
		}
	}
	return nil
}

func (l *MyList) Delete(x interface{}) {
	prev := l.FindPrevious(x)
	if prev != nil && !l.IsLast(prev) {
		prev.Next = prev.Next.Next
	}
}

func (l *MyList) FindPrevious(x interface{}) *Node {
	for cur := l.head; cur.Next != nil; cur = cur.Next {
		if cur.Next.Element == x {
			return cur
		}
	}
	return nil
}

func (l *MyList) Insert(x interface{}, position *Node) {
	if position != nil {
		position.Next = &Node{
			Element: x,
			Next:    position.Next,
		}
	}
}

func (l *MyList) DeleteList() {
	l.head.Next = nil
}

func (l *MyList) Header() *Node {
	return l.head
}

func (l *MyList) First() *Node {
	return l.head.Next
}

func (l *MyList) Advance(position *Node) *Node {
	return position.Next
}

func (l *MyList) Retrieve(position *Node) interface{} {
	return position.Element
}

package linkedlist

import (
	"github.com/hisoka-111/go-data-structures/list"
)

type LinkedList struct {
	head *list.Node
}

func NewLinkedList() list.List {
	return &LinkedList{
		head: &list.Node{},
	}
}

func (l *LinkedList) IsEmpty() bool {
	return l.head.Next == nil
}

func (l *LinkedList) IsLast(position *list.Node) bool {
	if position == nil {
		return false
	}
	return position.Next == nil
}

func (l *LinkedList) Find(x interface{}) *list.Node {
	for cur := l.head.Next; cur != nil; cur = cur.Next {
		if cur.Element == x {
			return cur
		}
	}
	return nil
}

func (l *LinkedList) Delete(x interface{}) {
	prev := l.FindPrevious(x)
	if prev != nil && !l.IsLast(prev) {
		prev.Next = prev.Next.Next
	}
}

func (l *LinkedList) FindPrevious(x interface{}) *list.Node {
	for cur := l.head; cur.Next != nil; cur = cur.Next {
		if cur.Next.Element == x {
			return cur
		}
	}
	return nil
}

func (l *LinkedList) Insert(x interface{}, position *list.Node) {
	if position != nil {
		position.Next = &list.Node{
			Element: x,
			Next:    position.Next,
		}
	}
}

func (l *LinkedList) DeleteList() {
	l.head.Next = nil
}

func (l *LinkedList) Header() *list.Node {
	return l.head
}

func (l *LinkedList) First() *list.Node {
	return l.head.Next
}

func (l *LinkedList) Advance(position *list.Node) *list.Node {
	return position.Next
}

func (l *LinkedList) Retrieve(position *list.Node) interface{} {
	return position.Element
}

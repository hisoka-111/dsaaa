package linkedlist

import (
	"fmt"
	"testing"

	"github.com/hisoka-111/go-data-structures/list"
)

func PrintList(l list.List) {
	fmt.Print("0")
	for cur := l.First(); cur != nil; cur = cur.Next {
		fmt.Printf(" -> %v", cur.Element)
	}
	fmt.Print("\n")
}

func TestLinkedList(t *testing.T) {
	l := NewLinkedList()
	l.Insert(1, l.Header())
	l.Insert(2, l.Header())
	PrintList(l)
	l.Insert(3, l.Find(2))
	PrintList(l)
	l.Delete(2)
	PrintList(l)
	l.DeleteList()
	PrintList(l)
}

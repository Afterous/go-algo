package SingleLinkedList

import (
	"testing"
)

func TestSingleLinkedList(t *testing.T){

	dll := NewSingleLinkedList()
	node := NewNode(1)
    dll.Append(node)
	node = NewNode(2)
	dll.Append(node)
	node = NewNode(3)
	dll.Append(node)
	node = NewNode(4)
	dll.Append(node)
	node = NewNode(5)
	dll.Append(node)
	dll.Prepend(NewNode(0))
    dll.Dumpit()
}
package DoubleLinkedList

import (
	nativeList "container/list"
	"fmt"
	"testing"
)

func TestDoubleLinkedList(t *testing.T){

	dll := NewDoubleLinkedList()
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

// "containers/list" is a double linked list
// see https://golang.org/pkg/container/list/
func TestNative(t *testing.T){
	l := nativeList.New()
    l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)
	l.PushFront(0)
	printnativeList(l)


}

func printnativeList(l *nativeList.List){
	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("Now the other way")
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
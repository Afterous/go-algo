package SingleLinkedList

import "fmt"

type node struct {
	value int
	next *node
}

type list struct {
	head *node
	tail *node
}

func NewSingleLinkedList() *list {
	return &list{
		nil,
		nil,
	}
}

func NewNode(value int) *node {
	return &node{value,nil}
}

func (l *list) Append(n *node){
	if l.head == nil {
		l.head = n
		l.tail = n
	} else{
		l.tail.next = n
		l.tail = n

	}
}

func (l *list) Prepend(n *node){
	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		n.next = l.head
		l.head = n
	}
}

func (l *list) Dumpit(){
	start := l.head
	for start != nil {
		fmt.Printf("%d\n",start.value)
		start = start.next
	}
}

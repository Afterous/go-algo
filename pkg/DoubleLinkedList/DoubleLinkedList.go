package DoubleLinkedList

import "fmt"

type node struct {
	value int
	nextNode *node // What is next (new is nil)
	lastNode *node // What was before.
}

type list struct {
	head *node // Points to start
	tail *node // Points to end
}

func NewDoubleLinkedList() *list {
	return &list{
		nil,
		nil,
	}
}

func NewNode(value int) *node {
	return &node{value,nil,nil}
}

func (l *list) Prepend(n *node){
	// If this is a new list.
	if l.head == nil {
		l.head = n
		l.tail = n
		// point to itself
		n.lastNode = n
		n.nextNode = n
		// Handling existing
	} else{
		l.head.lastNode = n
		n.nextNode = l.head
		l.head = n
	}
}

func (l *list) Append(n *node){
	// If this is a new list
	if l.tail == nil {
		l.head = n
		l.tail = n
		// Handling existing.
	} else {
		l.tail.nextNode = n
		n.lastNode = l.tail
		l.tail = n
	}
}

func (l *list) Dumpit(){
	start := l.head
	// Forwards
	for start != nil {
		fmt.Printf("%d\n",start.value)
		start = start.nextNode
	}
	fmt.Println("Now the other way")
	end := l.tail
	// Backwords
	for end != nil {
		fmt.Printf("%d\n",end.value)
		end = end.lastNode
	}
}
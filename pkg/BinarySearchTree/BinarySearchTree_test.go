package BinarySearchTree

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestSize(t *testing.T) {
	bree := Newbtree()
	elements := []int{3, 1, 4, 6, 9, 2, 5, 7}
	for _,v := range elements {
		n := NewNode(v)
		bree.BSTInsert(n)
	}
	assert.Equal(t,SizeOfbtree(bree),len(elements))
}


func TestPostOrderTravel(t *testing.T) {
	bree := Newbtree()
	elements := []int{3, 1, 4, 6, 9, 2, 5, 7}
	for _,v := range elements {
		n := NewNode(v)
		bree.BSTInsert(n)
	}
	n,_ := bree.Root()
	var buf bytes.Buffer
	log.SetFlags(0) // Don't log timestamp
	log.SetOutput(&buf) // write to buf instead.
	bree.PostOrderTravel(n)
	log.SetOutput(os.Stdout) // Revert change.
	assert.Equal(t,"2 \n1 \n5 \n7 \n9 \n6 \n4 \n3 \n",buf.String())

}

func TestPreOrderTravel(t *testing.T) {
	bree := Newbtree()
	elements := []int{3, 1, 4, 6, 9, 2, 5, 7}
	for _,v := range elements {
		n := NewNode(v)
		bree.BSTInsert(n)
	}
	n,_ := bree.Root()
	var buf bytes.Buffer
	log.SetFlags(0) // Don't log timestamp
	log.SetOutput(&buf) // write to buf instead.
	bree.PreOrderTravel(n)
	log.SetOutput(os.Stdout) // Revert change.
	assert.Equal(t,"3 \n1 \n2 \n4 \n6 \n5 \n9 \n7 \n",buf.String())
}

func TestNodeLevelNotFound(t *testing.T){
	bree := Newbtree()
	elements := []int{3, 1, 4, 6, 9, 2, 5, 7}
	for _,v := range elements {
		n := NewNode(v)
		bree.BSTInsert(n)
	}
	a:= NewNode(0)
	assert.Equal(t,-1,bree.NodeLevel(a))
}


func TestNodeLevelFound(t *testing.T){
	bree := Newbtree()
	elements := []int{3, 1, 4, 6, 9, 2, 5, 7}
	for _,v := range elements {
		n := NewNode(v)
		bree.BSTInsert(n)
	}
	a:= NewNode(9)
	assert.Equal(t,4,bree.NodeLevel(a))
}
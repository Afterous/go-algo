package main

import (
	"fmt"
	b "github.com/afterous/go-algo/pkg/BinarySearchTree"
)

func compare(btree *b.Btree,slice []int) bool {
	// Size check
	if b.SizeOfbtree(btree) != len(slice){
		return false
	}
	// iterate through slice, put value and search
	for _,v := range slice {
		_, err := btree.BSTSearch(v)
		if err !=nil {
			return false
		}

	}
	return true
}

func main(){
	fmt.Println("Hello CSC335, Binary tree")
	bree := b.Newbtree()
	elements := []int{3, 1, 4, 6, 9, 2, 5, 7}
	for _,v := range elements {
		n := b.NewNode(v)
		bree.BSTInsert(n)
	}
	a := compare(bree,elements)
	fmt.Println("They are the same: ",a)
}


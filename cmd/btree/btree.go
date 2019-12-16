package main

import (
	"fmt"
	"github.com/afterous/go-algo/pkg/BinarySearchTree"
)
func main() {
	fmt.Println("Hello CSC335, Binary tree")
	bree := BinarySearchTree.Newbtree()
	elements := []int{3, 1, 4, 6, 9, 2, 5, 7}
	//elements := []int{1, -10, 4,6,3,-11,-6,-7,-5}

	fmt.Println("Input:")
	for _,v := range elements {
		fmt.Printf("%d ",v)
		n := BinarySearchTree.NewNode(v)
		bree.BSTInsert(n)
	}
	fmt.Print("\n")
	n, _ := bree.Root()
	fmt.Printf("Root key: %d\n",n.Key())
	fmt.Printf("Size Of tree: %d\n",BinarySearchTree.SizeOfbtree(bree))
	fmt.Println("InOrder:")
	bree.InOrderTravel(n)
	fmt.Print("\n")
	fmt.Println("PreOrder:")
	bree.PreOrderTravel(n)
	fmt.Print("\n")
	fmt.Println("PostOrder:")
	bree.PostOrderTravel(n)
	fmt.Print("\n")
	fmt.Printf("REMOVING Root key: %d\n",n.Key())
	bree.BSTRemove(3)
	fmt.Printf("Size Of tree: %d\n",BinarySearchTree.SizeOfbtree(bree))
	fmt.Println("InOrder:")
	bree.InOrderTravel(n)
	fmt.Print("\n")
	fmt.Println("PreOrder:")
	bree.PreOrderTravel(n)
	fmt.Print("\n")
	fmt.Println("PostOrder:")
	bree.PostOrderTravel(n)
	fmt.Print("\n")
}
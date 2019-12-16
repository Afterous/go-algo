package BinarySearchTree

import (
	"fmt"
	"log"
)

type Btree struct {
	root *node
}

type node struct {
	key   int
	left  *node // key > lefNode.key
	right *node // key < leftNode.key
}

// Public method to get the size of a tree
func SizeOfbtree(tree *Btree) int {
	rnode, err := tree.Root()
	if err != nil {
		return 0 // size of zero
	}
	// Walk the nodes while keeping a counter
	// Couldn't use anonymous function to do something like:
	/*cnt := func up(n *node) uint32{ return up(n.left)+1+up(n.right)}(rnode) */
	count := size(rnode)
	return int(count)
}

// private, called by public method to get tree
func size(n *node) uint32 {
	if n == nil {
		return 0
	}
	return size(n.left) + 1 + size(n.right)
}

func (n *node) Key() int {
	return n.key
}

func (n *node) Min() int {
	if n.left == nil {
		return n.key
	}
	return n.left.Min()
}

func (n *node) Max() int {
	if n.right == nil {
		return n.key
	}
	return n.right.Max()
}
func Newbtree() *Btree {
	return &Btree{nil}
}

func NewNode(v int) *node {
	return &node{v, nil, nil}
}

// Public method to be given a node and determine if that node's key is present, if it isn't return -1 else return the height of the that node.
func (b *Btree) NodeLevel(n *node) int {
	// TODO: Fixme, we shouldn't "need" to call search, should we merge or track depth?
	_, err:=b.BSTSearch(n.key)
	if err != nil {
		return -1
	}
	return descend(b.root, n.key, 1)
}

// Private method to  descend into a tree to determine level.
func descend(n *node, data int, lvl int) int {
	if n == nil {
		return 0
	} else if n.key == data {
		return lvl
	}
	dlvl := descend(n.left,data,lvl+1)
	if dlvl != 0 {
		return dlvl
	}
	dlvl = descend(n.right,data,lvl+1)
	return dlvl
}

func (b *Btree) Root() (*node, error) {
	if b.root != nil {
		return b.root, nil
	}
	return nil, fmt.Errorf("root node not found for binary tree")
}

func (b *Btree) InOrderTravel(node *node) {
	if node != nil {
		b.InOrderTravel(node.left)
		log.Printf("%d ", node.key)
		b.InOrderTravel(node.right)
	}
}

func (b *Btree) PreOrderTravel(node *node) {
	if node != nil {
		log.Printf("%d ", node.key)
		b.PreOrderTravel(node.left)
		b.PreOrderTravel(node.right)
	}
}

func (b *Btree) PostOrderTravel(node *node) {
	if node != nil {
		b.PostOrderTravel(node.left)
		b.PostOrderTravel(node.right)
		log.Printf("%d ", node.key)
	}
}

func (b *Btree) BSTRemove(key int) {
	var parent *node       // Need to track the parent of the node to be deleted
	current, _ := b.Root() // current will point to the node to be deleted
	//find the node to be removed
	for current != nil && current.key != key {
		if key < current.key {
			parent = current
			current = current.left
		} else {
			parent = current
			current = current.right
		}
	}

	if current != nil { // The node to be deleted is found
		//Case A : the node is a leaf node
		if current.left == nil && current.right == nil {
			if parent.right == current {
				parent.right = nil
			} else {
				parent.left = nil
				current = &node{}
			}
			//Case B: the node has two children
			// Must find the smallest key in the right subtree of the node, which is
			//found by going to the right child of the node then all the way to the left.
		} else if current.left != nil && current.right != nil {
			succ := current.right // go to the right child of the node to be removed
			parent = current      // initialize parent node
			if succ.left == nil { // right child of the node has no left child
				parent.right = succ.right
				current.key = succ.key
				succ = &node{} // delete
			} else {                   //otherwise keep going left
				for succ.left != nil { // then find the smallest key in the left subtree
					parent = succ
					succ = succ.left
				}
				current.key = succ.key //Replace the node to be deleted by the succ node
				parent.left = nil      // skip the succ node
				succ = &node{}         // delete
			}
		} else {                   // Case C: Node has one child
			if b.root.key == key { //if the node is the root node treat differently
				fmt.Println("HERE")
				if b.root.right != nil {
					b.root = b.root.right
				} else {
					b.root = b.root.left
				}
			} else // a non-root node with one child to be removed
			{
				if current.left != nil {
					parent.left = current.left
				} else {
					parent.right = current.right
				}
			}
			current = &node{} // delete
		}
	}
}

// Binary Tree Search, given a key will return point of Node that matched, else error.
func (b *Btree) BSTSearch(key int) (*node, error) {
	cur := b.root
	for cur != nil {
		if key == cur.key {
			return cur, nil // found
		} else if key < cur.key {
			cur = cur.left
		} else {
			cur = cur.right
		}

	}
	return nil, fmt.Errorf("key: %d not found in b tree", key)
}

// Binary Search Tree Insert, given a node
func (b *Btree) BSTInsert(node *node) {
	// empty tree lets add
	if b.root == nil {
		b.root = node
		node.left = nil
		node.right = nil
		// We have a root, check where to add
	} else {
		cur := b.root
		for cur != nil {
			if node.key < cur.key {
				if cur.left == nil {
					cur.left = node
					cur = nil
				} else {
					cur = cur.left
				}
			} else {
				if cur.right == nil {
					cur.right = node
					cur = nil
				} else {
					cur = cur.right
				}
			}
		}
		node.left = nil
		node.right = nil
	}
}

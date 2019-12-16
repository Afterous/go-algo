package AVLTree

type Avl interface {
	heightUpdate(n *avlnode)
	RotateRight(n *avlnode)
	RotateLeft(n *avlnode)
	Rebalance(n *avlnode)
	Insert(n avlnode)
	Remove(i int)
	inOrderTravel(n *avlnode)

}

func (tree *AVLtree) heightUpdate(n *avlnode) {
	var lf int = -1
	var rh int = -1

	if n.left != nil {
		lf =  n.left.nodeHeight
	}
	if n.right != nil {
		rh = n.right.nodeHeight
	}
	ans	:= func(lh,rh int) int {
		if lh > rh {
			return lh+1
		} else {
			return rh+1
		}
	}(lf,rh)
	n.nodeHeight = ans
}

func (tree *AVLtree) RotateRight(n *avlnode) {
	panic("implement me")
}

func (tree *AVLtree) RotateLeft(n *avlnode) {
	panic("implement me")
}

func (tree *AVLtree) Rebalance(n *avlnode) {
	panic("implement me")
}

func (tree *AVLtree) Insert(n avlnode) {
	panic("implement me")
}

func (tree *AVLtree) Remove(i int) {
	panic("implement me")
}

func (tree *AVLtree) inOrderTravel(n *avlnode) {
	panic("implement me")
}

type avlnode struct {
	parent,right,left *avlnode
	element,nodeHeight int
}

type AVLtree struct {
	root *avlnode
}

func NewAVLtree() *AVLtree {
	return &AVLtree{}
}


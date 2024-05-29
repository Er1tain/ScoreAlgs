package main

type TreeNode struct {
	value       int
	left_child  *TreeNode
	right_child *TreeNode
}

func NewTreeNode(value int) *TreeNode {

	return &TreeNode{value: value, left_child: nil, right_child: nil}
}

func (tree_node *TreeNode) setRightChild(child_node *TreeNode) {
	tree_node.right_child = child_node
}

func (tree_node *TreeNode) setLeftChild(child_node *TreeNode) {
	tree_node.left_child = child_node
}

func main() {

}

package main

import "fmt"

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

func (tree_node *TreeNode) print() {
	fmt.Println(tree_node.value)
}

func (tree_node *TreeNode) getLeftChild() *TreeNode {
	return tree_node.left_child
}

func (tree_node *TreeNode) getRightChild() *TreeNode {
	return tree_node.right_child
}

func main() {

	root := NewTreeNode(8)
	root.setLeftChild(NewTreeNode(5))
	root.setRightChild(NewTreeNode(13))

	root.getRightChild().print()

}

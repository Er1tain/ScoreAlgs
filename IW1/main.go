package main

import (
	"fmt"
)

type TreeNode struct {
	value       int
	left_child  *TreeNode
	right_child *TreeNode
}

func CreateNodeTree(value int) *TreeNode {

	return &TreeNode{value: value, left_child: nil, right_child: nil}
}

func (tree_node *TreeNode) setRightChild(child_node *TreeNode) {
	tree_node.right_child = child_node
}

func (tree_node *TreeNode) setLeftChild(child_node *TreeNode) {
	tree_node.left_child = child_node
}

func addNode(root *TreeNode, value int) {
	if root == nil {
		return
	}
	if value <= root.value {
		if root.left_child == nil {
			root.setLeftChild(CreateNodeTree(value))
		} else {
			addNode(root.left_child, value)
		}
	} else {
		if root.right_child == nil {
			root.setRightChild(CreateNodeTree(value))
		} else {
			addNode(root.right_child, value)
		}
	}
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

func createBinarryTree(root *TreeNode, values []int) {
	for _, value := range values {
		addNode(root, value)
	}
}

func countNodesWithOneChild(root *TreeNode) int {
	if root == nil {
		return 0
	}

	count := 0
	if (root.getLeftChild() != nil && root.getRightChild() == nil) || (root.getLeftChild() == nil && root.getRightChild() != nil) {
		count++
	}

	count += countNodesWithOneChild(root.getLeftChild())
	count += countNodesWithOneChild(root.getRightChild())

	return count
}

func main() {
	// values := []int{8, 1, 3, 5, 7, 2}
	// values := []int{4, -5, 7, 12, -3, 9}
	//values := []int{4, 2, 5, -3, 3, 5, 7, 6, 8}
	values := []int{4, 2, 5, -3, 3, 7, 6, 8}

	root := CreateNodeTree(values[0])
	values = values[1:]
	createBinarryTree(root, values)

	count_onechild_nodes := countNodesWithOneChild(root)
	fmt.Println(count_onechild_nodes)

}

package leetcode

import (
	"github.com/yskang/leetcodeUtil/leetData"
)

// Flatten is a solution of a problem '114. Flatten Binary Tree to Linked List' of leetcode.
func Flatten(root *leetData.TreeNode) {
	flatten(root)
}

func flatten(root *leetData.TreeNode) {
	if root == nil {
		return
	}
	p := root
	for node := range nodeGen(root) {
		if node == p {
			continue
		}
		p.Left = nil
		p.Right = node
		p = node
	}
}

func nodeGen(root *leetData.TreeNode) <-chan *leetData.TreeNode {
	c := make(chan *leetData.TreeNode)

	go func() {
		var walk func(r *leetData.TreeNode)
		walk = func(r *leetData.TreeNode) {
			if r == nil {
				return
			}
			left, right := r.Left, r.Right
			c <- r
			walk(left)
			walk(right)
		}
		walk(root)
		close(c)
	}()

	return c
}

// use generator and list
func flatten__(root *leetData.TreeNode) {
	if root == nil {
		return
	}
	nodes := make([]*leetData.TreeNode, 0)
	for node := range getNode(root) {
		nodes = append(nodes, node)
	}

	for _, node := range nodes[1:] {
		root.Left = nil
		root.Right = node
		root = root.Right
	}
}

func getNode(root *leetData.TreeNode) <-chan *leetData.TreeNode {
	c := make(chan *leetData.TreeNode)

	go func() {
		var fr func(r *leetData.TreeNode)
		fr = func(r *leetData.TreeNode) {
			if r == nil {
				return
			}
			c <- r
			fr(r.Left)
			fr(r.Right)
		}
		fr(root)
		close(c)
	}()

	return c
}

// recursive way
func flatten_(root *leetData.TreeNode) {
	flattenTree(root)
}

func flattenTree(root *leetData.TreeNode) *leetData.TreeNode {
	if root == nil {
		return nil
	}

	if root.Left != nil && root.Right != nil {
		tmp := root.Right
		root.Right = root.Left
		root.Left = nil
		last := flattenTree(root.Right)
		last.Right = tmp
		return flattenTree(tmp)
	} else if root.Left != nil {
		root.Right = root.Left
		root.Left = nil
		return flattenTree(root.Right)
	} else if root.Right != nil {
		return flattenTree(root.Right)
	}
	return root
}

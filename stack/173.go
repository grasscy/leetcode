package stack

import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 二叉搜索树的一个重要的特性是是二叉搜索树的中序序列是升序序列；因此，中序遍历是该解决方案的核心。

type BSTIterator struct {
	s []int
}

func Constructor(root *TreeNode) BSTIterator {
	traversal := preorderTraversal(root)
	sort.Ints(traversal)
	return BSTIterator{s: traversal}
}

func (this *BSTIterator) Next() int {
	r := this.s[0]
	this.s = this.s[1:]
	return r
}

func (this *BSTIterator) HasNext() bool {
	return len(this.s) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

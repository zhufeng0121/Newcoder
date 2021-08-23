package main

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 *
 * @param root TreeNode类 the root of binary tree
 * @return int整型二维数组
 */
func threeOrders( root *TreeNode ) [][]int {
	res := make([][]int,0)
	res = append(res,preorder(root),inorder(root),postorder(root))
	return res

	// write code here
}

func preorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int,0)
	res = append(res, root.Val)
	res = append(res, preorder(root.Left)...)
	res = append(res, preorder(root.Right)...)
	return res
}
func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int,0)
	res = append(res, inorder(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorder(root.Right)...)
	return res
}

func postorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int,0)
	res = append(res, postorder(root.Left)...)
	res = append(res, postorder(root.Right)...)
	res = append(res, root.Val)
	return res
}


func preorderI(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int,0)
	if root == nil {
		return nil
	}
	stack = append(stack,root)
	for len(stack) != 0 {
		cur := stack[len(stack) -1]
		stack = stack[:len(stack) - 1]
		//先对右子树进行入栈
		res = append(res, cur.Val)
		if cur.Right != nil {
			stack = append(stack,cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack,cur.Left)
		}

	}
	return res

}






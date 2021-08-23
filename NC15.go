package main

type TreeNode struct {
	Val      int
	Left     *TreeNode
	Right    *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode,0)
	var result [][]int
	queue = append(queue, root)
	for len(queue) != 0{
		level := make([]int,0)
		size := len(queue)
		for i := 0; i < size; i++ {
			temp := queue[0]
			queue = queue[1:]
			if temp.Left != nil {
				queue = append(queue,temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue,temp.Right)
			}
			level = append(level, temp.Val)
		}
		result = append(result, level)
	}
	return result

}

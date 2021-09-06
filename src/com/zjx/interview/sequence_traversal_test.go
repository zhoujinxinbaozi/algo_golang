/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/3/1 3:12 下午
 */

/**
    层序遍历二叉树
 */
package interview

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 层序遍历 获取每一层的最大值
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var resultList []int
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) != 0 {
		tmpQueue := queue
		queue = []*TreeNode{}
		result := math.MinInt32
		for len(tmpQueue) != 0 {
			node := tmpQueue[0]
			result = max(result, node.value)
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
			tmpQueue = tmpQueue[1:]
		}
		resultList = append(resultList, result)
	}
	return resultList
}
/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/10/14 3:03 下午
 */

package _13_pathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
var result [][]int
func pathSum(root *TreeNode, sum int) [][]int {
	if root != nil {
		dfs(root, 0, sum, []int{})
	}
	return result
}

func dfs(root *TreeNode, target, sum int, arr []int) {
	arr = append(arr, root.Val)
	if target+root.Val == sum && root.Left == nil && root.Right == nil {
		var tmp []int
		tmp = make([]int, len(arr))
		copy(tmp, arr)
		result = append(result, tmp)
		return
	}

	if root.Left != nil {
		dfs(root.Left, target+root.Val, sum, arr)
	}

	if root.Right != nil {
		dfs(root.Right, target+root.Val, sum, arr)
	}
	arr = arr[:len(arr)-1]
}
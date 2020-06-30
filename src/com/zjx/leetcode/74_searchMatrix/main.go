package _4_searchMatrix

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/27 08:42
 */

/**
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
示例 1:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
输出: true
示例 2:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
输出: false
*/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	begin := 0
	end := len(matrix)*len(matrix[0]) - 1

	for begin <= end {
		mid := (begin + end) / 2
		if matrix[mid/len(matrix[0])][mid%len(matrix[0])] == target { // 中心位置计算方法
			return true
		} else if matrix[mid/len(matrix[0])][mid%len(matrix[0])] < target {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}

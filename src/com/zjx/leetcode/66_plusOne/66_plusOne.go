package _6_plusOne

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/26 17:10
 */

/**
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。

*/
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			continue
		}
		digits[i] = digits[i] + 1 // 数组中有非9的数字，加1返回
		return digits
	}
	result := make([]int, len(digits)+1) // 数组中所有数字都为9，长度加1
	result[0] = 1                        // 首位置1
	return result
}

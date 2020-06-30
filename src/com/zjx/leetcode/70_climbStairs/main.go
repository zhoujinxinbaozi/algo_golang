package _0_climbStairs

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/26 20:18
 */

/**
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。
*/
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	i1 := 1
	i2 := 2
	for i := 3; i <= n; i++ {
		i1, i2 = i2, i1+i2
	}
	return i2
}

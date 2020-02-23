package main

import "fmt"

/**
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

solution:
	回溯算法即可

 */
func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	var ans []string
	backtrack(&ans, "", 0, 0, n)
	return ans
}

/**
	ans:全局结果集
	cur:当前获得的字符串
	open:左括号的数量
	close：右括号的数量
	max：括号的对数
 */
func backtrack(ans *[]string, cur string, open int, close int, max int) {
	if len(cur) == max*2 {
		*ans = append(*ans, cur)
	}
	if open < max {
		backtrack(ans, cur+"(", open+1, close, max)
	}
	if close < open {
		backtrack(ans, cur+")", open, close+1, max)
	}
}

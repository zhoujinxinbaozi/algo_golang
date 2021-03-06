package main

import "fmt"

/**
 * @Auther: zhoujinxin@bytedance.com
 * @Date: 2020/6/26 20:33
 */

/**
给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符


示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')
*/
func main() {
	minDistance("sea", "eat")
}

func minDistance(word1 string, word2 string) int {
	if len(word2) == 0 {
		return len(word1)
	}
	if len(word1) == 0 {
		return len(word2)
	}

	result := make([][]int, len(word1))
	for index, _ := range result {
		result[index] = make([]int, len(word2))
	}
	if word2[0] != word1[0] {
		result[0][0] = 1
	}
	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			if i == 0 && j == 0 { // init
				continue
			}
			if i == 0 {
				if word1[i] == word2[j] {
					result[i][j] = GetMin(result[i][j-1]+1, j)
				} else {
					result[i][j] = result[i][j-1] + 1
				}
				continue
			}
			if j == 0 {
				if word1[i] == word2[j] {
					result[i][j] = GetMin(result[i-1][j]+1, i)
				} else {
					result[i][j] = result[i-1][j] + 1
				}
				continue
			}
			if word1[i] == word2[j] {
				result[i][j] = GetMin(GetMin(result[i-1][j-1], result[i-1][j]+1), result[i][j-1]+1)
				continue
			}
			result[i][j] = GetMin(GetMin(result[i-1][j-1]+1, result[i-1][j]+1), result[i][j-1]+1)
		}
	}
	fmt.Println(result)
	return result[len(word1)-1][len(word2)-1]
}

func GetMin(i, j int) int {
	if i < j {
		return i
	}
	return j
}

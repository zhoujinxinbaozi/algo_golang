/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/3/1 3:12 下午
 */

/**
    层序遍历二叉树
 */
package interview

import "fmt"

// Tree
type Tree struct {
	v int
	l *Tree
	r *Tree
}

// sequenceTraversal 层序遍历 按行进行打印
func sequenceTraversal(root *Tree) {
	if root == nil {
		return
	}

	var nodeList []*Tree
	nodeList = append(nodeList, root)

	mark := root

	for len(nodeList) != 0 {
		tmp := nodeList[0]
		fmt.Print(tmp.v, "\t")
		var cur *Tree
		if tmp.l != nil {
			nodeList = append(nodeList, tmp.l)
			cur = tmp.l
		}
		if tmp.r != nil {
			nodeList = append(nodeList, tmp.r)
			cur = tmp.r
		}
		if tmp == mark {
			fmt.Println()
			mark = cur
		}
		nodeList = nodeList[1:]
	}
}
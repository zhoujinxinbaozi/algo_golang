/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/3/1 3:22 下午
 */

package interview

import (
	"fmt"
	"strconv"
	"testing"
)

// isMirTree 是否是镜像二叉树
func isMirTree(root *Tree) bool {
	if root == nil {
		return true
	}

	return isMir(root.l, root.r)
}

// isMir 左右子树是否是镜像
func isMir(l, r *Tree) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	return l.v == r.v && isMir(l.l, r.r) && isMir(l.r, r.l)
}

// treeMir 二叉树的镜像二叉树
func treeMir(root *Tree) *Tree {
	if root == nil {
		return nil
	}
	root.l, root.r = treeMir(root.r), treeMir(root.l)
	return root
}

func TestSprintf(t *testing.T) {
	s := fmt.Sprintf("%5s", strconv.Itoa(12))
	fmt.Println(s)
}
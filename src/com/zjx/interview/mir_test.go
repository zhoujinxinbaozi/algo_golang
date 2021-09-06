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
func isMirTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMir(root.left, root.right)
}

// isMir 左右子树是否是镜像
func isMir(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	return l.value == r.value && isMir(l.left, r.right) && isMir(l.right, r.left)
}

// treeMir 二叉树的镜像二叉树
func treeMir(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.left, root.right = treeMir(root.right), treeMir(root.left)
	return root
}

func TestSprintf(t *testing.T) {
	s := fmt.Sprintf("%5s", strconv.Itoa(12))
	fmt.Println(s)
}

func TestChange(t *testing.T) {
	arr := []int{3,4}
	fmt.Printf("test change addr : %p\n", arr)
	fmt.Printf("%v\n", arr)
	change1(arr)
	//change(&arr)
	fmt.Printf("%v\n", arr)
}

func change(arr *[]int) {
	(*arr)[0], (*arr)[1] = (*arr)[1], (*arr)[0]
	*arr = append(*arr, 5)
	fmt.Printf("change addr : %p\n", *arr)
}

func change1(arr []int) {
	fmt.Printf("change1 addr : %p\n", arr)
}
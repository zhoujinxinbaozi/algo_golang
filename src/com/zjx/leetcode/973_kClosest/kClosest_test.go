/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/11/14 9:57 上午
 */

package _73_kClosest

import (
	"sort"
)

type Point struct {
	X int
	Y int
}

func kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		p, q := points[i], points[j]
		return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
	})
	return points[:k]
}
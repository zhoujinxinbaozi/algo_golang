/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/6/1 3:06 下午
 */

package interview

import (
	"fmt"
	"sort"
	"testing"
)

func heapsort(nums []int) []int {
	lenght := len(nums)
	if lenght == 0 {
		return nil
	}
	for i := 0; i < lenght; i++ {
		// 第一下 使得第一个元素最小，接下来就从第二个来构造，使得下一个最小，
		miniHeap(nums[i:])
	}
	return nums
}

// 使得每次操作完毕 堆顶的元素就是最小的，由于堆的特性，我们只需要从倒数第二层开始就可以了
func miniHeap(nums []int) {
	length := len(nums)
	floor := length/2 - 1
	for i := floor; i >= 0; i-- {
		// 然后比较的每一个节点与其两个孩子节点的大小，使得爸爸永远是最小的
		// 有一种特殊情况，就是最后一个节点的孩子节点可能不存在，和可能只有一个，所以需要加上一个判断
		father := i
		left := 2*i + 1
		right := 2*i + 2

		if right < length && nums[father] > nums[right] {
			nums[father], nums[right] = nums[right], nums[father]
		}

		if left < length && nums[father] > nums[left] {
			nums[father], nums[left] = nums[left], nums[father]
		}
	}
}

func TestTopN(t *testing.T) {
	topN(3)
}

func topN(n int) {

	bigData := []int{
		3,7,6,2,4,5,1,
	}
	fmt.Println("the big data is ", bigData)
	startNums := bigData[:n]
	miniHeap(startNums)
	for _, data := range bigData[n:] {
		if data <= startNums[0] {
			continue
		} else {
			startNums[0] = data
			miniHeap(startNums)
		}
	}
	fmt.Println("top10 is ", heapsort(startNums))
}

func TestThreeSum(t *testing.T) {
	fmt.Printf("%v\n", threeSum([]int{0,0,0}))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	resultList := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i-1 >= 0 && nums[i] == nums[i-1]{
			continue
		}
		j := i+1
		k := len(nums) - 1
		for j < k {
			if nums[i] + nums[j] + nums[k] == 0 {
				resultList = append(resultList, []int{nums[i], nums[j], nums[k]})
				fmt.Println(i, j, k)
				j++
				for j -1 >= 0 && j < len(nums) && nums[j] == nums[j-1] {
					j++
				}
				k--
				for k>= 0 && k+1 < len(nums) && nums[k+1] == nums[k] {
					k--
				}
			} else if nums[i] + nums[j] + nums[k] < 0 {
				j++
				for j-1 >= 0 && j < len(nums) && nums[j] == nums[j-1] {
					j++
				}
			} else {
				k--
				for k >= 0 && k+1 < len(nums) && nums[k+1] == nums[k] {
					k--
				}
			}
		}
	}
	return resultList
}
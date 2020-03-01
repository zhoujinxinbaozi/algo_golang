package main

/**
	problem:
	给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
	不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
	solution:

 */
func main() {

}

func removeElement(nums []int, val int) int {
	begin := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[begin] = nums[i]
			begin++
		}
	}
	return begin
}

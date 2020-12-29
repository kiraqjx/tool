package go1

func twoSum(nums []int, target int) []int {
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		for j := i + 1; j < numsLen; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return make([]int, 0)
}

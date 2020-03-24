package basic

func twoSum(nums []int, target int) []int {
	l := len(nums)
	var r []int
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if target == nums[i] + nums[j] {
				return append(r, i, j)
			}
		}
	}
	return r
}
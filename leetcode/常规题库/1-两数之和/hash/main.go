package hash


func twoSum(nums []int, target int) []int {
	l := len(nums)
	var r []int
	t := make(map[int]int)
	for i := 0; i < l; i++ {
		t[target - nums[i]] = i
		if i + 1 < l {
			if e, ok := t[nums[i + 1]]; ok {
				return append(r, i + 1, e)
			}
		}
	}
	return r
}


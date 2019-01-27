package t1_two_sum

// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	for idx, v := range nums {
		for i := idx+1; i < len(nums); i++ {
			if v + nums[i] == target {
				return []int{idx, i}
			}
		}
	}
	return []int{}
}


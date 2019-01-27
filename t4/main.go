package main

// https://leetcode.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if 0 == l1 && 0 == l2 {return 0}
	if 0 == l1 { return findMedianForOneArr(nums2) }
	if 0 == l2 { return findMedianForOneArr(nums1) }

	l := l1 + l2
	des := 0
	if l % 2 == 1 {
		des = l/2
	} else {
		des = l/2 - 1
	}
	median := 0
	curNum := 0
	i, j := 0, 0
	for i < l1 && j < l2 {
		if curNum == des {
			break
		}
		if nums1[i] < nums2[j] {
			median = nums1[i]
			i++
		} else {
			median = nums1[j]
			j++
		}
		curNum ++
	}
	if i >= l1 {
		for j < l2 {
			if curNum == des {
				if l % 2 == 1 {
					return float64(median)
				}
				break
			}
			if nums1[i] < nums2[j] {
				median = nums1[i]
				i++
			} else {
				median = nums1[j]
				j++
			}
		}

	}
	if j >= l2 {

	}

}

func findMedianForOneArr(nums []int) (median float64) {
	l := len(nums)
	if l % 2 == 1 {
		return float64(nums[l/2])
	} else {
		return float64(nums[l/2]+nums[l/2-1])/2
	}
}

package main

import (
	"fmt"
	"math"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	bMap := map[uint8]int{}
	bArr := []byte(s)
	maxLen := 0
	for i, v := range bArr {
		if bMap[v] == 0 {
			bMap[v] = i
		} else {
			currentLen := getCurrentLen(bMap)
			if getCurrentLen(bMap) > maxLen {
				maxLen = currentLen
			}
			resetMap(bMap, v)
			bMap[v] = i
		}
	}
	currentLen := getCurrentLen(bMap)
	if currentLen > maxLen {
		maxLen = currentLen
	}
	return maxLen
}

func resetMap(bMap map[uint8]int, k uint8)  {
	v := bMap[k]
	for i, j := range bMap {
		if j < v {
			delete(bMap, i)
		}
	}
}

func getCurrentLen(bMap map[uint8]int) (len int) {
	min := math.MaxInt64
	max := 0
	for _, v := range bMap {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min + 1
}

func main() {
	fmt.Println(lengthOfLongestSubstring(" "))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
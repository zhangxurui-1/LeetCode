// https://leetcode.cn/problems/two-sum/description/?envType=study-plan-v2&envId=top-100-liked

package code

import (
	"fmt"
	"sort"
)

type element struct {
	loc   int
	value int
}

func twoSum(nums []int, target int) []int {
	var ans = []int{}
	var nums2 = []element{}
	for i := 0; i < len(nums); i++ {
		nums2 = append(nums2, element{loc: i, value: nums[i]})
	}
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i].value < nums2[j].value
	})

	var (
		i int = 0
		j int = len(nums2) - 1
	)

	for i < j {
		if nums2[i].value+nums2[j].value == target {
			ans = append(ans, nums2[i].loc, nums2[j].loc)
			return ans
		} else if nums2[i].value+nums2[j].value < target {
			i++
		} else {
			j--
		}
	}
	return ans
}

func Test() {
	/*
		var (
			nums       = []int{2, 7, 11, 15}
			target int = 9
		)
	*/

	var (
		nums       = []int{3, 2, 4}
		target int = 6
	)

	fmt.Println(twoSum(nums, target))
}

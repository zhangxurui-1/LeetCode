package code

import (
	"fmt"
	"sort"
)

/*
func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var (
		i   int = 0
		j   int = len(nums) - 1
		k   int = (i + j) / 2
		tmp int = 0
		vi  int = 0
		vj  int = 0
		ans     = [][]int{}
	)
	fmt.Println(nums)
	for i < len(nums) && nums[i] <= 0 {
		vi = nums[i]
		j = len(nums) - 1
		for j > i && nums[j] >= 0 {
			vj = nums[j]
			k = (i + j) / 2
			tmp = -(vi + vj)
			if tmp > vj || tmp < vi {
				for ; j > 0 && nums[j] == vj; j-- {
				}
				continue
			}
			if nums[k] == tmp && i < k && k < j {
				ans = append(ans, []int{vi, tmp, vj})
				// fmt.Println(i, "\t", k, "\t", j)
			} else if nums[k] < tmp {
				for ; k < j && nums[k] < tmp; k++ {
				}
				if nums[k] == tmp && i < k && k < j {
					ans = append(ans, []int{vi, tmp, vj})
					// fmt.Println(i, "\t", k, "\t", j)
				}
			} else {
				for ; k > i && nums[k] > tmp; k-- {
				}
				if nums[k] == tmp && i < k && k < j {
					ans = append(ans, []int{vi, tmp, vj})
					// fmt.Println(i, "\t", k, "\t", j)
				}
			}

			for ; j > 0 && nums[j] == vj; j-- {
			}
		}
		for ; i < len(nums) && nums[i] == vi; i++ {
		}
	}
	return ans

}
*/

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var (
		i   int = 0
		j   int = 0
		k   int = 0
		vi  int = 0
		vj  int = 0
		vk  int = 0
		ans     = [][]int{}
	)

	fmt.Println(nums)
	for i < len(nums) && nums[i] <= 0 {
		vi = nums[i]
		j = i + 1
		k = len(nums) - 1
		for j < k {
			vj = nums[j]
			vk = nums[k]
			if vj+vk == -vi && i < j && j < k {
				ans = append(ans, []int{vi, vj, vk})
				for ; j < len(nums) && nums[j] == vj; j++ {
				}
				for ; k > 0 && nums[k] == vk; k-- {
				}
			} else if vj+vk < (-vi) {
				for ; j < len(nums) && nums[j] == vj; j++ {
				}
			} else {
				for ; k > 0 && nums[k] == vk; k-- {
				}
			}
		}
		for ; i < len(nums) && nums[i] == vi; i++ {
		}
	}
	return ans

}

func Test15() {
	var nums = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

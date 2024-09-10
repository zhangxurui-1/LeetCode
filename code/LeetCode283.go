package code

import "fmt"

func moveZeroes(nums []int) {
	var (
		i int = 0
		j int = 0
	)
	for j < len(nums) {
		if nums[i] == 0 {
			for ; j < len(nums) && nums[j] == 0; j++ {
			}
			if j >= len(nums) {
				break
			}
			nums[i] = nums[j]
			nums[j] = 0
		} else {
			nums[i] = nums[j]
		}
		i++
		j++
	}
	for ; i < len(nums); i++ {
		nums[i] = 0
	}
}

func Test283() {
	var nums = []int{0}
	moveZeroes(nums)
	fmt.Println(nums)
}

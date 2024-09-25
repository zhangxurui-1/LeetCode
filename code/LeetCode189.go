package code

import "fmt"

func swap(nums []int, i, j int) {
	nums[i] = nums[i] + nums[j]
	nums[j] = nums[i] - nums[j]
	nums[i] = nums[i] - nums[j]
}

/*
// 原地 环状替换
func gcd(x, y int) int {
	for y != 0 {
		x = x % y
		// swap
		x = x + y
		y = x - y
		x = x - y
	}
	return x
}

func rotate(nums []int, k int) {
	steps := k % len(nums)
	if steps == 0 {
		return
	}

	for i := 0; i < gcd(len(nums), steps); i++ {
		j := i + steps
		for j != i {
			swap(nums, i, j)
			j = (j + steps) % len(nums)
		}
	}
}
*/

// 数组翻转
func flip(nums []int, i, j int) {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func rotate(nums []int, k int) {
	steps := k % len(nums)
	if steps == 0 {
		return
	}
	flip(nums, 0, len(nums)-1)
	flip(nums, 0, steps-1)
	flip(nums, steps, len(nums)-1)
}

func Test189() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

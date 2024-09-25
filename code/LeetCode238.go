package code

import "fmt"

func productExceptSelf(nums []int) []int {
	var (
		ans = make([]int, len(nums))
	)
	// left-->right
	ans[0] = 1
	for i := 1; i < len(ans); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}

	// right-->left
	tmp := 1
	for i := len(ans) - 2; i >= 0; i-- {
		tmp = tmp * nums[i+1]
		ans[i] = ans[i] * tmp
	}
	return ans
}

func Test238() {
	var nums = []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
}

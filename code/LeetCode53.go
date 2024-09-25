package code

import "fmt"

/*
func maxSubArray(nums []int) int {
	var (
		maxValue int = nums[0]
		tmp      int = 0
	)

	for _, v := range nums {
		tmp += v
		maxValue = max(maxValue, tmp)
		if tmp <= 0 {
			tmp = 0
		}
	}

	return maxValue
}
*/

// 分治
type Info struct {
	lmax     int // 包含线段左端点的最大子段和
	rmax     int // 包含线段右端点的最大子段和
	maxValue int // 最大子段和
	sum      int // 区间和
}

func getMaxValue(nums []int, i int, j int) Info {
	if i == j {
		return Info{nums[i], nums[i], nums[i], nums[i]}
	}
	m := (i + j) / 2
	linfo := getMaxValue(nums, i, m)
	rinfo := getMaxValue(nums, m+1, j)
	var info Info
	info.lmax = max(linfo.lmax, linfo.sum+rinfo.lmax)
	info.rmax = max(rinfo.rmax, rinfo.sum+linfo.rmax)
	info.maxValue = max(linfo.maxValue, rinfo.maxValue, linfo.rmax+rinfo.lmax)
	info.sum = linfo.sum + rinfo.sum
	return info
}

func maxSubArray(nums []int) int {
	info := getMaxValue(nums, 0, len(nums)-1)
	return info.maxValue
}

func Test53() {
	var nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

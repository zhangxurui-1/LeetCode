// https://leetcode.cn/problems/longest-consecutive-sequence/?envType=study-plan-v2&envId=top-100-liked

package code

import "fmt"

func longestConsecutive(nums []int) int {
	var (
		m       = make(map[int]uint8)
		ans int = 0
		cnt int = 0
	)
	for _, n := range nums {
		m[n] = 1
	}
	for _, n := range nums {
		if _, ok := m[n-1]; !ok && m[n] != 2 {
			cnt = 0
			for j := n; m[j] == 1; j++ {
				cnt++
			}

			ans = max(ans, cnt)
			m[n] = 2
		}
	}
	return ans
}

func Test128() {
	var m = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(m))
}

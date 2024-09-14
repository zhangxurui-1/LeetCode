package code

import "fmt"

// O(N^2)
/*
func subarraySum(nums []int, k int) int {
	var (
		sum int = 0
		ans int = 0
	)
	for i, _ := range nums {
		sum = 0
		for _, v2 := range nums[i:] {
			sum += v2
			if sum == k {
				ans++
			}
		}
	}
	return ans
}
*/

// 前缀和 + 哈希表优化
/*
func subarraySum(nums []int, k int) int {
	var (
		sum int = 0
		vis     = make(map[int][]int)
		ans int = 0
	)

	for i, v := range nums {
		sum += v
		if _, exist := vis[sum]; !exist {
			vis[sum] = []int{i}
		} else {
			vis[sum] = append(vis[sum], i)
		}
	}

	for key, value := range vis {
		if key == k {
			ans += len(value)
		}

		indices, ok := vis[key-k]
		if !ok {
			continue
		} else {
			for _, v := range value {
				for _, index := range indices {
					if index < v {
						ans++
					}
				}
			}
		}
	}
	return ans
}
*/

// 再优化一点
func subarraySum(nums []int, k int) int {
	var (
		mp      = make(map[int]int, 0)
		ans     = 0
		pre int = 0
	)

	for _, v := range nums {
		pre += v
		if pre == k {
			ans++
		}
		if value, exist := mp[pre-k]; exist {
			ans += value
		}

		// 更新mp
		if _, exist := mp[pre]; !exist {
			mp[pre] = 1
		} else {
			mp[pre]++
		}
	}

	return ans
}

func Test560() {
	var nums = []int{1, -1, 0}
	fmt.Println(subarraySum(nums, 0))
}

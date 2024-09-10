package code

import "fmt"

// time limit
/*
func maxArea(height []int) int {
	var ans int = 0
	for i, h := range height {
		for j := i + 1; j < len(height); j++ {
			ans = max(ans, (j-i)*min(height[j], h))
		}
	}
	return ans
}
*/

func maxArea(height []int) int {
	var (
		ans int = 0
		i   int = 0
		j   int = len(height) - 1
	)

	ans = (j - i) * min(height[i], height[j])
	for i < j {
		l := height[i]
		r := height[j]
		if l <= r {
			for ; i < len(height) && height[i] <= l; i++ {
			}
			if i >= j {
				break
			}
			ans = max(ans, min(height[i], r)*(j-i))
		} else {
			for ; j >= 0 && height[j] <= r; j-- {
			}
			if i >= j {
				break
			}
			ans = max(ans, min(l, height[j])*(j-i))
		}
	}
	return ans

}

func Test11() {
	var (
		height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	)
	fmt.Println(maxArea(height))
}

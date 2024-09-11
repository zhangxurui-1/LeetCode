package code

import "fmt"

func trap(height []int) int {
	var (
		i   int = 0
		j   int = 0
		k   int = 0
		ans int = 0
	)
	for i < len(height)-1 {
		for ; i < len(height)-1 && height[i] <= height[i+1]; i++ {
		}
		j = i + 1
		for ; j < len(height)-1 && height[j] >= height[j+1]; j++ {
		}
		k = j + 1
		for ; k < len(height)-1 && height[k] <= height[k+1]; k++ {
		}

		// fmt.Println(i, " ", j, " ", k)
		if k < len(height) {
			var tmp int = min(height[i], height[k])
			for l := i; l <= k; l++ {
				ans += (tmp - min(height[l], tmp))
			}
		}

		i = k
	}

	return ans
}

func Test42() {
	var height = []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(height))
}
